CREATE TABLE IF NOT EXISTS customer (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_full_name` varchar(30) NOT NULL,
  `customer_email` varchar(50) NOT NULL,
  `customer_phone_number` varchar(15) NOT NULL,
  `customer_username` varchar(30) NOT NULL,
  `customer_password` varchar(30) NOT NULL,
  `account_status` int(1) NOT NULL,
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS menu (
  `menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(100) NOT NULL,
  `price` float NOT NULL,
  `menu_type_id` int(11) NOT NULL,
  `ingredients` varchar(500) NOT NULL,
  `menu_status` int(1) NOT NULL,
  PRIMARY KEY (`menu_id`),
  KEY `menu_type_id` (`menu_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS menu_type (
  `menu_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `type_name` varchar(50) NOT NULL,
  `description` varchar(100) NOT NULL,
  PRIMARY KEY (`menu_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS orders (
  `order_id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `order_date` date NOT NULL,
  `total_price` float,
  `order_status` int(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`order_id`),
  KEY `customer_id` (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS order_details (
  `order_details_id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `menu_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `total_price` float,
  PRIMARY KEY (`order_details_id`),
  KEY `order_id` (`order_id`,`menu_id`),
  KEY `menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS payment (
  `payment_id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `amount` float NOT NULL,
  `payment_type` varchar(50) NOT NULL,
  `payment_date` date NOT NULL,
  PRIMARY KEY (`payment_id`),
  KEY `order_id` (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

CREATE TABLE IF NOT EXISTS rating (
  `rating_id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_id` int(11) NOT NULL,
  `score` int(1) NOT NULL,
  `remarks` varchar(100) NOT NULL,
  `date_recorded` date NOT NULL,
  `customer_id` int(11) NOT NULL,
  PRIMARY KEY (`rating_id`),
  KEY `menu_id` (`menu_id`),
  KEY `customer_id` (`customer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

ALTER TABLE menu
  ADD CONSTRAINT `tblmenu_ibfk_1` FOREIGN KEY (`menu_type_id`) REFERENCES menu_type (`menu_type_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE orders
  ADD CONSTRAINT `tblorder_ibfk_2` FOREIGN KEY (`customer_id`) REFERENCES customer (`customer_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE order_details
  ADD CONSTRAINT `tblorderdetails_ibfk_2` FOREIGN KEY (`order_id`) REFERENCES orders (`order_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `tblorderdetails_ibfk_1` FOREIGN KEY (`menu_id`) REFERENCES menu (`menu_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE payment
  ADD CONSTRAINT `tblpayment_ibfk_2` FOREIGN KEY (`order_id`) REFERENCES orders (`order_id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE rating
  ADD CONSTRAINT `tblrating_ibfk_2` FOREIGN KEY (`customer_id`) REFERENCES customer (`customer_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `tblrating_ibfk_1` FOREIGN KEY (`menu_id`) REFERENCES menu (`menu_id`) ON DELETE CASCADE ON UPDATE CASCADE;
