#!/usr/bin/env Rscript

suppressPackageStartupMessages(library(dplyr))
suppressPackageStartupMessages(library(lubridate))

args = commandArgs(trailingOnly = TRUE)
message(sprintf("The trial is %s\nthe date is %s\nthe int file is %s\nthe string file is %s", args[1L], args[2L], args[3L], args[4L]))

screening <- read.csv(args[3L], stringsAsFactors = F)
screening_txt <- read.csv(args[4L], stringsAsFactors = F)
screening <- merge(screening, screening_txt, by = "ResponseId")
analyze <- screening %>%
  mutate(StartDate.x = date(StartDate.x) ) %>%
  arrange(StartDate.x) %>%
  filter(StartDate.x >= date(args[2L]) ) %>%
  filter(Progress.x >= 99 & (!is.na(PSS12.x) & PSS12.x != "") ) %>%
  mutate(KTEN.score.x = KTEN.score.x + 10) %>%
  filter(KTEN.score.x >= 20) %>%
  filter(SIDAS.score.x < 21) %>%
  select(ResponseId, CTD1.x, CTD2a.x, CTD2b.x, CTD3.x, Q356.y) %>%
  rename(Qualtrics_ID = ResponseId, name = CTD1.x, phone = CTD2a.x, phone_confirmed = CTD2b.x, email = CTD3.x, consent = Q356.y)

write.csv(analyze, paste0("Screening_data_",args[1L],"_", Sys.Date(),".csv"))

message(sprintf("The number of the eligible participants is %d", nrow(analyze)))

