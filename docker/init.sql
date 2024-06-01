-- init.sql

-- Create the RssScraper database
CREATE DATABASE "RssScraper";

-- Connect to the RssScraper database and create a table
\connect "RssScraper"

-- Create a table in the RssScraper database
CREATE TABLE example_table (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Insert some data into the example_table
INSERT INTO example_table (name) VALUES ('Sample data');
