CREATE DATABASE IF NOT EXISTS moneyFlow;

USE moneyFlow;

CREATE TABLE IF NOT EXISTS m_bop_categories(
    m_bop_category_id int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    bop_name varchar(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS t_users(
    user_id int  UNSIGNED AUTO_INCREMENT,
    user_name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS t_bops(
    t_bop_id int UNSIGNED AUTO_INCREMENT PRIMARY KEY ,
    payment_name varchar(255) NOT NULL,
    payment_date DATE NOT NULL,
    total_amount  int NOT NULL,
    m_bop_category_id int UNSIGNED NOT NULL, 
    user_id int UNSIGNED NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (m_bop_category_id) REFERENCES m_bop_categories(m_bop_category_id) ON DELETE CASCADE,
    FOREIGN KEY t_bops (user_id) REFERENCES t_users (user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bop_details(
    bop_detail_id int UNSIGNED AUTO_INCREMENT,
    purchase_num varchar(255) NOT NULL,
    amount  int NOT NULL,
     t_bop_id int UNSIGNED NOT NULL ,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (bop_detail_id),
    FOREIGN KEY(t_bop_id) REFERENCES t_bops(t_bop_id) ON DELETE CASCADE
);
