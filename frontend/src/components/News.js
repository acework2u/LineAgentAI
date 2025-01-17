import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

const newsData = [
    {
        id: 1,
        title: 'New Medical Volunteer Program Launches',
        description: 'Join our latest medical volunteer program to help communities in need.',
        image: 'https://via.placeholder.com/300x200',
        link: '#',
    },
    {
        id: 2,
        title: 'Health Awareness Seminar Announced',
        description: 'Learn about the importance of regular health check-ups in our upcoming seminar.',
        image: 'https://via.placeholder.com/300x200',
        link: '#',
    },
    {
        id: 3,
        title: 'Fundraising Event for Medical Supplies',
        description: 'Support our initiative to provide medical supplies to underprivileged areas.',
        image: 'https://via.placeholder.com/300x200',
        link: '#',
    },
];

const News = () => {
    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">News ประชาสัมพันธ์</h2>
            <div className="row gy-4">
                {newsData.map((news) => (
                    <div className="col-12 col-md-6 col-lg-4" key={news.id}>
                        <div className="card h-100 shadow-sm">
                            <img src={news.image} className="card-img-top" alt={news.title} />
                            <div className="card-body">
                                <h5 className="card-title">{news.title}</h5>
                                <p className="card-text">
                                    {news.description.length > 100
                                        ? `${news.description.substring(0, 100)}...`
                                        : news.description}
                                </p>
                                <a href={news.link} className="btn btn-primary w-100">
                                    Read More
                                </a>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default News;