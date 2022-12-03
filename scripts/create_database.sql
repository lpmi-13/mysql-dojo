CREATE DATABASE IF NOT EXISTS transactions;

CREATE TABLE IF NOT EXISTS `transactions`.`referrals` (
    referral_id mediumint NOT NULL AUTO_INCREMENT,
    currency varchar(255),
    email varchar(255),
    referral_code varchar(255),
    referral_submitted varchar(255),
    referree varchar(255),
    referrer varchar(255),
    username varchar(255),
    PRIMARY KEY (referral_id)
    );