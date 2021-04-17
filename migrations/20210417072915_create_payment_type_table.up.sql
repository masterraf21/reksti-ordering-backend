CREATE TABLE IF NOT EXISTS payment_type (
    `payment_type_id` int(11) NOT NULL AUTO_INCREMENT,
    `payment_method` VARCHAR(50) NOT NULL,
    `payment_company` VARCHAR(50) NOT NULL,
    PRIMARY KEY(`payment_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

ALTER TABLE payment
    ADD CONSTRAINT `tblpayment_1bfk_1` FOREIGN KEY (`payment_type_id`) REFERENCES payment_type (`payment_type_id`) ON DELETE CASCADE ON UPDATE CASCADE; 