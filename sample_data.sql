-- Sample data for hotels database

-- Insert Managers
INSERT INTO managers (name, about, phone, created_at, updated_at)
VALUES 
  ('Priya Sharma', 'Hospitality professional with 15 years of experience in luxury properties', '+919876543210', NOW(), NOW()),
  ('Raj Patel', 'Former chef turned hotel manager with a passion for guest satisfaction', '+918765432109', NOW(), NOW()),
  ('Sarah Johnson', 'International hotel management expert specializing in boutique hotels', '+16175551234', NOW(), NOW()),
  ('Amit Verma', 'Award-winning manager with background in both business and hospitality', '+917654321098', NOW(), NOW()),
  ('Michael Chen', 'Experienced manager focused on sustainable and eco-friendly hospitality', '+6590123456', NOW(), NOW());

-- Insert Hotels
INSERT INTO hotels (
  name, location_short, full_address, pin_code, landmark, price, rating, 
  image, availability, description, location_link, superhost, host_story,
  cancellation_policy, checkin_time, manager_id, amenities, created_at, updated_at
)
VALUES
  (
    'Sea View Resort', 
    'Mumbai, Maharashtra', 
    '123 Marine Drive, Nariman Point, Mumbai', 
    '400021', 
    'Near Gateway of India', 
    8500.00, 
    4.7, 
    'https://example.com/images/seaview-resort.jpg', 
    true, 
    'Luxurious seaside resort with panoramic views of the Arabian Sea. Features spacious rooms, multiple dining options, and a world-class spa.',
    'https://maps.google.com/?q=18.9633,72.8180', 
    true, 
    'Started in 2005 as a small guest house and grew into one of the most recognized properties in Mumbai.',
    'Free cancellation 24 hours before check-in. 50% refund up to 12 hours before check-in.',
    '2023-01-01 14:00:00', 
    1, 
    '["WiFi", "Pool", "Spa", "Room Service", "Restaurant", "Bar", "Gym", "Beach Access", "Free Parking", "Airport Shuttle"]',
    NOW(), 
    NOW()
  ),
  
  (
    'Mountain Retreat', 
    'Shimla, Himachal Pradesh', 
    '45 Mall Road, Shimla', 
    '171001', 
    'Near Christ Church', 
    6200.00, 
    4.5, 
    'https://example.com/images/mountain-retreat.jpg', 
    true, 
    'Charming mountain hotel with colonial architecture and breathtaking views of the Himalayas. Cozy rooms with fireplaces and an excellent restaurant serving local cuisine.',
    'https://maps.google.com/?q=31.1048,77.1734', 
    true, 
    'A 100-year-old heritage building carefully restored to maintain its historical charm while offering modern amenities.',
    'Free cancellation up to 48 hours before check-in. 25% penalty afterward.',
    '2023-01-01 12:00:00', 
    2, 
    '["WiFi", "Restaurant", "Room Service", "Fireplace", "Mountain Views", "Tea Garden", "Library", "Board Games", "Hiking Trails"]',
    NOW(), 
    NOW()
  ),
  
  (
    'Urban Oasis Hotel', 
    'Bangalore, Karnataka', 
    '78 MG Road, Bangalore', 
    '560001', 
    'Near Cubbon Park', 
    7800.00, 
    4.8, 
    'https://example.com/images/urban-oasis.jpg', 
    true, 
    'Modern luxury hotel in the heart of Bangalore\'s tech district. Features smart rooms, co-working spaces, and a rooftop infinity pool with city views.',
    'https://maps.google.com/?q=12.9716,77.5946', 
    false, 
    'Created by a team of tech entrepreneurs to offer the perfect blend of work and relaxation spaces.',
    'Free cancellation up to 24 hours. Non-refundable afterward.',
    '2023-01-01 15:00:00', 
    3, 
    '["WiFi", "Pool", "Rooftop Bar", "Co-working Space", "Gym", "Restaurant", "Smart Home Features", "EV Charging", "Bike Rental"]',
    NOW(), 
    NOW()
  ),
  
  (
    'Royal Heritage Haveli', 
    'Jaipur, Rajasthan', 
    '156 Amber Road, Jaipur', 
    '302002', 
    'Near Amber Fort', 
    12500.00, 
    4.9, 
    'https://example.com/images/royal-heritage.jpg', 
    true, 
    'Restored 18th-century royal hunting lodge turned into a luxury boutique hotel. Features hand-painted frescoes, private courtyards, and traditional Rajasthani architecture.',
    'https://maps.google.com/?q=26.9124,75.7873', 
    true, 
    'Owned by descendants of the royal family who have preserved the property\'s rich history while adding modern luxuries.',
    'Free cancellation up to 7 days before check-in. 50% penalty up to 72 hours. Non-refundable afterward.',
    '2023-01-01 14:00:00', 
    4, 
    '["WiFi", "Pool", "Heritage Tours", "Cooking Classes", "Spa", "Restaurant", "Bar", "Garden", "Cultural Performances"]',
    NOW(), 
    NOW()
  ),
  
  (
    'Backwater Bliss Resort', 
    'Alleppey, Kerala', 
    '23 Punnamada, Kuttanad, Alleppey', 
    '688006', 
    'Punnamada Lake', 
    9300.00, 
    4.6, 
    'https://example.com/images/backwater-bliss.jpg', 
    true, 
    'Eco-friendly resort with private houseboats and lake-view villas. Experience Kerala\'s famous backwaters with traditional architecture and Ayurvedic treatments.',
    'https://maps.google.com/?q=9.4981,76.3388', 
    true, 
    'Started by a local family passionate about showcasing Kerala\'s natural beauty while preserving its ecosystem.',
    'Free cancellation up to 48 hours. 30% penalty afterward.',
    '2023-01-01 13:00:00', 
    5, 
    '["WiFi", "Private Houseboat", "Ayurvedic Spa", "Restaurant", "Yoga Classes", "Fishing", "Cooking Classes", "Bird Watching", "Organic Farm"]',
    NOW(), 
    NOW()
  );

-- Insert Rooms
INSERT INTO rooms (hotel_id, room_type, bed_type, capacity, price, available, created_at, updated_at)
VALUES
  -- Rooms for Sea View Resort (hotel_id = 1)
  (1, 'Standard Room', 'Queen', 2, 8500.00, true, NOW(), NOW()),
  (1, 'Deluxe Room', 'King', 2, 12000.00, true, NOW(), NOW()),
  (1, 'Sea View Suite', 'King', 3, 18000.00, true, NOW(), NOW()),
  (1, 'Family Room', 'Twin and Queen', 4, 15000.00, false, NOW(), NOW()),

  -- Rooms for Mountain Retreat (hotel_id = 2)
  (2, 'Heritage Room', 'Queen', 2, 6200.00, true, NOW(), NOW()),
  (2, 'Mountain View Room', 'King', 2, 8500.00, true, NOW(), NOW()),
  (2, 'Luxury Suite', 'King', 2, 12000.00, true, NOW(), NOW()),
  (2, 'Family Cottage', 'King and Twin', 4, 15000.00, true, NOW(), NOW()),

  -- Rooms for Urban Oasis Hotel (hotel_id = 3)
  (3, 'Smart Room', 'Queen', 2, 7800.00, true, NOW(), NOW()),
  (3, 'Executive Room', 'King', 2, 9500.00, true, NOW(), NOW()),
  (3, 'City View Suite', 'King', 2, 14000.00, true, NOW(), NOW()),
  (3, 'Business Suite', 'King', 3, 16500.00, false, NOW(), NOW()),

  -- Rooms for Royal Heritage Haveli (hotel_id = 4)
  (4, 'Heritage Room', 'Queen', 2, 12500.00, true, NOW(), NOW()),
  (4, 'Royal Suite', 'King', 2, 18000.00, true, NOW(), NOW()),
  (4, 'Maharaja Suite', 'King', 3, 25000.00, true, NOW(), NOW()),
  (4, 'Royal Family Suite', 'King and Twin', 4, 32000.00, true, NOW(), NOW()),

  -- Rooms for Backwater Bliss Resort (hotel_id = 5)
  (5, 'Garden Villa', 'Queen', 2, 9300.00, true, NOW(), NOW()),
  (5, 'Lake View Villa', 'King', 2, 12500.00, true, NOW(), NOW()),
  (5, 'Houseboat Suite', 'Queen', 2, 15000.00, true, NOW(), NOW()),
  (5, 'Premium Houseboat', 'King', 3, 18500.00, false, NOW(), NOW()); 