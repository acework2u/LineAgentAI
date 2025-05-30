import React, {useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import {Table} from 'react-bootstrap'
import axios from "axios";

const API_URL = "http://localhost:8081/api/v1/events";

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



const EventManagement = () => {
    // application config
    const [events, setEvents] = React.useState([]);

    //Mount
    useEffect(() => {
        fetchEvents();
    })

    //Fetch events
    const fetchEvents = async () => {
        try {
            const response = await axios.get(API_URL);
            setEvents(response.data.events);
            console.log(response.data.events);
        } catch (error) {
            console.log("Error fetching events",error);
        }
    }


    // render
    return (
        <>
        <div className="container mt-5">
            <div className="row justify-content-center align-items-center">Event Management</div>
        <div className={"row"}>
            <div className={"col-2"}>
                <h3>Left menu</h3>

            </div>
            <div className={"col-10"}>
                <Table striped bordered hover className={"mt-3-3"}>
                    <thead>
                    <tr>
                        <th>Event Name</th>
                        <th>Date</th>
                        <th>Description</th>
                    </tr>
                    </thead>
                    <tbody>
                    {events.map((event) => (
                        <tr key={event.id}>
                            <td>{event.title}</td>
                            <td>{event.startDate}</td>
                            <td>{event.description}</td>
                        </tr>
                    ))}
                    </tbody>
                </Table>


            </div>
        </div>



        </div>
        </>
    );
}
export default EventManagement;