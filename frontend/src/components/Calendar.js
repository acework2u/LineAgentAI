import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

const events = [
    {
        id: 1,
        date: '2025-01-20',
        title: 'Community Health Check-Up',
        description: 'Provide free medical services to the local community.',
    },
    {
        id: 2,
        date: '2025-02-10',
        title: 'Free Vaccination Drive',
        description: 'Join us in offering free vaccinations to underserved populations.',
    },
    {
        id: 3,
        date: '2025-03-05',
        title: 'Dental Care Outreach',
        description: 'Help us provide dental care to those in need.',
    },
];

const Calendar = () => {
    const daysInMonth = 31; // Example: For January
    const startDayOfWeek = 1; // Example: Monday (0 = Sunday, 1 = Monday, etc.)

    const getEventsForDate = (date) => {
        return events.filter((event) => event.date === date);
    };

    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">Calendar ปฏิทินงาน</h2>
            <p className="text-center">
                Explore upcoming events for this month. Click on a date to view details.
            </p>

            {/* Calendar Grid */}
            <div className="table-responsive">
                <table className="table table-bordered text-center">
                    <thead>
                    <tr>
                        <th>Sun</th>
                        <th>Mon</th>
                        <th>Tue</th>
                        <th>Wed</th>
                        <th>Thu</th>
                        <th>Fri</th>
                        <th>Sat</th>
                    </tr>
                    </thead>
                    <tbody>
                    {/* Calendar Rows */}
                    {Array.from({ length: 6 }).map((_, weekIndex) => (
                        <tr key={weekIndex}>
                            {Array.from({ length: 7 }).map((_, dayIndex) => {
                                const dayNumber = weekIndex * 7 + dayIndex - startDayOfWeek + 1;

                                // Check if day is valid (falls within the current month)
                                if (dayNumber < 1 || dayNumber > daysInMonth) {
                                    return <td key={dayIndex}></td>;
                                }

                                // Check if there are events for this date
                                const date = `2025-01-${String(dayNumber).padStart(2, '0')}`;
                                const dayEvents = getEventsForDate(date);

                                return (
                                    <td key={dayIndex} className={dayEvents.length > 0 ? 'bg-light' : ''}>
                                        <div>
                                            <strong>{dayNumber}</strong>
                                        </div>
                                        {dayEvents.map((event) => (
                                            <small key={event.id} className="d-block text-primary">
                                                {event.title}
                                            </small>
                                        ))}
                                    </td>
                                );
                            })}
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>

            {/* Event Details */}
            <div className="mt-4">
                <h4>Upcoming Events</h4>
                {events.map((event) => (
                    <div key={event.id} className="mb-3 p-3 border rounded">
                        <h5>{event.title}</h5>
                        <p>
                            <strong>Date:</strong> {event.date}
                        </p>
                        <p>{event.description}</p>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Calendar;