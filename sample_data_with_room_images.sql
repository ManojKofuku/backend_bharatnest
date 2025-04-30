-- Sample data for hotels database with multiple room images

-- Insert Managers
INSERT INTO managers (name, about, phone, created_at, updated_at)
VALUES 
  ('Priya Sharma', 'Hospitality professional with 15 years of experience in luxury properties', '+919876543210', NOW(), NOW()),
  ('Raj Patel', 'Former chef turned hotel manager with a passion for guest satisfaction', '+918765432109', NOW(), NOW()),
  ('Sarah Johnson', 'International hotel management expert specializing in boutique hotels', '+16175551234', NOW(), NOW());

-- Insert Hotels (3 hotels example)
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
    'https://example.com/images/hotels/seaview-resort.jpg', 
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
    'https://example.com/images/hotels/mountain-retreat.jpg', 
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
    'https://example.com/images/hotels/urban-oasis.jpg', 
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
  );

-- Insert Rooms with multiple images (3 rooms per hotel)
INSERT INTO rooms (
  hotel_id, room_type, bed_type, capacity, price, available, description, 
  images, created_at, updated_at
)
VALUES
  -- Rooms for Sea View Resort (hotel_id = 1)
  (
    1, 'Standard Room', 'Queen', 2, 8500.00, true, 
    'Comfortable room with modern amenities and partial sea view. Includes a queen-sized bed, work desk, and ensuite bathroom.',
    '["https://example.com/images/rooms/seaview-standard-1.jpg", "https://example.com/images/rooms/seaview-standard-2.jpg", "https://example.com/images/rooms/seaview-standard-3.jpg"]', 
    NOW(), NOW()
  ),
  (
    1, 'Deluxe Room', 'King', 2, 12000.00, true, 
    'Spacious room with full sea view and private balcony. Features a king-sized bed, sitting area, and luxury bathroom with bathtub.',
    '["https://example.com/images/rooms/seaview-deluxe-1.jpg", "https://example.com/images/rooms/seaview-deluxe-2.jpg", "https://example.com/images/rooms/seaview-deluxe-3.jpg"]', 
    NOW(), NOW()
  ),
  (
    1, 'Sea View Suite', 'King', 3, 18000.00, true, 
    'Premium suite with panoramic sea views, separate living room, and private terrace. Includes king-sized bed, dining area, and marble bathroom with jacuzzi.',
    '["https://example.com/images/rooms/seaview-suite-1.jpg", "https://example.com/images/rooms/seaview-suite-2.jpg", "https://example.com/images/rooms/seaview-suite-3.jpg"]', 
    NOW(), NOW()
  ),

  -- Rooms for Mountain Retreat (hotel_id = 2)
  (
    2, 'Heritage Room', 'Queen', 2, 6200.00, true, 
    'Classic room with colonial decor and wooden furnishings. Features a queen-sized bed, writing desk, and views of the garden.',
    '["https://example.com/images/rooms/mountain-heritage-1.jpg", "https://example.com/images/rooms/mountain-heritage-2.jpg", "https://example.com/images/rooms/mountain-heritage-3.jpg"]', 
    NOW(), NOW()
  ),
  (
    2, 'Mountain View Room', 'King', 2, 8500.00, true, 
    'Charming room with stunning mountain views and a private balcony. Includes a king-sized bed, seating area with fireplace, and luxurious bathroom.',
    '["https://example.com/images/rooms/mountain-view-1.jpg", "https://example.com/images/rooms/mountain-view-2.jpg", "https://example.com/images/rooms/mountain-view-3.jpg"]', 
    NOW(), NOW()
  ),
  (
    2, 'Luxury Suite', 'King', 2, 12000.00, true, 
    'Elegant suite with separate living area and panoramic views of the Himalayas. Features antique furniture, fireplace, king-sized bed, and premium bathroom with clawfoot tub.',
    '["https://example.com/images/rooms/mountain-suite-1.jpg", "https://example.com/images/rooms/mountain-suite-2.jpg", "https://example.com/images/rooms/mountain-suite-3.jpg"]', 
    NOW(), NOW()
  ),

  -- Rooms for Urban Oasis Hotel (hotel_id = 3)
  (
    3, 'Smart Room', 'Queen', 2, 7800.00, true, 
    'Modern room with smart technology integration. Features voice-controlled lighting and temperature, queen-sized bed, and sleek bathroom with rain shower.',
    '["https://example.com/images/rooms/urban-smart-1.jpg", "https://example.com/images/rooms/urban-smart-2.jpg", "https://example.com/images/rooms/urban-smart-3.jpg"]', 
    NOW(), NOW()
  ),
  (
    3, 'Executive Room', 'King', 2, 9500.00, true, 
    'Business-friendly room with dedicated work area and city views. Includes king-sized bed, ergonomic chair, high-speed internet, and modern bathroom.',
    '["https://example.com/images/rooms/urban-executive-1.jpg", "https://example.com/images/rooms/urban-executive-2.jpg", "https://example.com/images/rooms/urban-executive-3.jpg"]', 
    NOW(), NOW()
  ),
  (
    3, 'City View Suite', 'King', 2, 14000.00, true, 
    'Spacious suite with floor-to-ceiling windows offering panoramic city views. Features a separate living area, king-sized bed, kitchenette, and luxury bathroom.',
    '["https://example.com/images/rooms/urban-suite-1.jpg", "https://example.com/images/rooms/urban-suite-2.jpg", "https://example.com/images/rooms/urban-suite-3.jpg"]', 
    NOW(), NOW()
  ); 