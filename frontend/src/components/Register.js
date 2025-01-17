import React, { useState } from 'react';

const Register = () => {
    // State for form inputs
    const [formData, setFormData] = useState({
        fullName: '',
        email: '',
        password: '',
    });

    // State for form submission message
    const [message, setMessage] = useState('');

    // Handle input changes
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    // Handle form submission
    const handleSubmit = (e) => {
        e.preventDefault();
        if (formData.fullName && formData.email && formData.password) {
            setMessage('Registration successful!');
            setFormData({ fullName: '', email: '', password: '' });
        } else {
            setMessage('Please fill out all fields.');
        }
    };

    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">Register สมัครสมาชิก</h2>
            <div className="row justify-content-center">
                <div className="col-md-6">
                    {/* Form */}
                    <form onSubmit={handleSubmit}>
                        {/* Full Name */}
                        <div className="mb-3">
                            <label htmlFor="fullName" className="form-label">Full Name</label>
                            <input
                                type="text"
                                id="fullName"
                                name="fullName"
                                className="form-control"
                                placeholder="Enter your full name"
                                value={formData.fullName}
                                onChange={handleChange}
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
                            />
                        </div>

                        {/* Password */}
                        <div className="mb-3">
                            <label htmlFor="password" className="form-label">Password</label>
                            <input
                                type="password"
                                id="password"
                                name="password"
                                className="form-control"
                                placeholder="Enter your password"
                                value={formData.password}
                                onChange={handleChange}
                            />
                        </div>

                        {/* Submit Button */}
                        <button type="submit" className="btn btn-primary w-100">
                            Register
                        </button>
                    </form>

                    {/* Submission Message */}
                    {message && (
                        <div className={`mt-3 alert ${message.includes('successful') ? 'alert-success' : 'alert-danger'}`}>
                            {message}
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default Register;