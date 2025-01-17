import React, { useState } from 'react';
import axios from 'axios';

const ContactUs = () => {
    // State for form inputs
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        subject: '',
        message: '',
    });

    // State for submission status
    const [statusMessage, setStatusMessage] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    // Handle input changes
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    // Handle form submission
    const handleSubmit = async (e) => {
        e.preventDefault();

        // Check if all fields are filled
        if (!formData.name || !formData.email || !formData.subject || !formData.message) {
            setStatusMessage('Please fill out all fields.');
            return;
        }

        try {
            setIsSubmitting(true);
            setStatusMessage('');

            // API request to send contact form data
            const response = await axios.post('http://api.acework.ai/contact-us', formData);

            if (response.status === 200) {
                setStatusMessage('Your message has been sent successfully!');
                setFormData({ name: '', email: '', subject: '', message: '' }); // Clear form
            }
        } catch (error) {
            if (error.response) {
                setStatusMessage(error.response.data.message || 'Failed to send your message. Please try again.');
            } else {
                setStatusMessage('An error occurred. Please try again later.');
            }
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">Contact Us ติดต่อ</h2>
            <p className="text-center mb-4">
                Have any questions or concerns? Please fill out the form below, and we will get back to you as soon as possible.
            </p>
            <div className="row justify-content-center">
                <div className="col-md-8">
                    {/* Contact Form */}
                    <form onSubmit={handleSubmit}>
                        {/* Name */}
                        <div className="mb-3">
                            <label htmlFor="name" className="form-label">Full Name</label>
                            <input
                                type="text"
                                id="name"
                                name="name"
                                className="form-control"
                                placeholder="Enter your full name"
                                value={formData.name}
                                onChange={handleChange}
                                required
                            />
                        </div>

                        {/* Email */}
                        <div className="mb-3">
                            <label htmlFor="email" className="form-label">Email</label>
                            <input
                                type="email"
                                id="email"
                                name="email"
                                className="form-control"
                                placeholder="Enter your email"
                                value={formData.email}
                                onChange={handleChange}
                                required
                            />
                        </div>

                        {/* Subject */}
                        <div className="mb-3">
                            <label htmlFor="subject" className="form-label">Subject</label>
                            <input
                                type="text"
                                id="subject"
                                name="subject"
                                className="form-control"
                                placeholder="Enter the subject"
                                value={formData.subject}
                                onChange={handleChange}
                                required
                            />
                        </div>

                        {/* Message */}
                        <div className="mb-3">
                            <label htmlFor="message" className="form-label">Message</label>
                            <textarea
                                id="message"
                                name="message"
                                className="form-control"
                                rows="5"
                                placeholder="Enter your message"
                                value={formData.message}
                                onChange={handleChange}
                                required
                            ></textarea>
                        </div>

                        {/* Submit Button */}
                        <button
                            type="submit"
                            className="btn btn-primary w-100"
                            disabled={isSubmitting}
                        >
                            {isSubmitting ? 'Sending...' : 'Send Message'}
                        </button>
                    </form>

                    {/* Status Message */}
                    {statusMessage && (
                        <div className={`mt-3 alert ${statusMessage.includes('successfully') ? 'alert-success' : 'alert-danger'}`}>
                            {statusMessage}
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default ContactUs;