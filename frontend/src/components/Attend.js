import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

const eventData = [
    {
        id: 1,
        title: 'Community Health Check-Up',
        date: 'January 20, 2025',
        location: 'Bangkok Community Center',
        description: 'Participate in our community health check-up event to provide free medical services.',
        image: 'https://via.placeholder.com/300x200',
    },
    {
        id: 2,
        title: 'Free Vaccination Drive',
        date: 'February 10, 2025',
        location: 'Chiang Mai Public Hall',
        description: 'Join us in offering free vaccinations to underserved populations.',
        image: 'https://via.placeholder.com/300x200',
    },
    {
        id: 3,
        title: 'Dental Care Outreach',
        date: 'March 5, 2025',
        location: 'Phuket Health Clinic',
        description: 'Help us provide dental care to those in need during this outreach program.',
        image: 'https://via.placeholder.com/300x200',
    },
];

const Attend = () => {
    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">Attend ร่วมงานแพทยอาสา</h2>
            <p className="text-center mb-4">
                Join our volunteer events to make a difference in the community. Browse upcoming events below and sign up to participate.
            </p>
            <div className="row gy-4">
                {eventData.map((event) => (
                    <div className="col-12 col-md-6 col-lg-4" key={event.id}>
                        <div className="card h-100 shadow-sm">
                            <img src={event.image} className="card-img-top" alt={event.title} />
                            <div className="card-body">
                                <h5 className="card-title">{event.title}</h5>
                                <p className="card-text">
                                    <strong>Date:</strong> {event.date}<br />
                                    <strong>Location:</strong> {event.location}
                                </p>
                                <p className="card-text">
                                    {event.description.length > 80
                                        ? `${event.description.substring(0, 80)}...`
                                        : event.description}
                                </p>
                                <button className="btn btn-primary w-100">Join Now</button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Attend;