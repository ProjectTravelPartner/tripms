USE tripms;
CREATE TABLE trips (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    accid INT,
    source VARCHAR(128),
    destination VARCHAR(128),
    travel_date DATE,
    vehicle VARCHAR(128),
    accomodation INT,
    costperperson VARCHAR(128)
)