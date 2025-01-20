import React, { useState } from 'react';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';

const Register = () => {

    const courses = [
        {title:"ปธพX",value:"ปธพX"},
        {title:"ปธพ1,2,3,4,5,6,7,8,9,10,11",value:"ปธพ1,2,3,4,5,6,7,8,9,10,11"},
        {title:"ปนพ1,2,3",value:"ปนพ1,2,3"},
        {title:"ปอพ1,2,3",value:"ปอพ1,2,3"},
        {title:"ปกพ1",value:"ปกพ1"},

    ]

    const CourseList = () =>{
       return courses.map((item,index)=> {
               return (
               <option key={index} value={item.value}>{item.title}</option>
       )

       })
    }


    const [formData, setFormData] = useState({
        phone: '',
        pin: '',
        firstName: '',
        lastName: '',
        role: '', // 'Med' or 'Non-Med'
        specialty: '',
        nonMedText: '',
        organize: '',
        position: '',
        dob: '',
        course: '',
        lineID: '',
        facebook: '',
        instagram: '',
        email: '',
        foodAllergy: '',
        religion: '',
        otherReligion: '',
        note: '',
    });

    const [isVerificationStep, setIsVerificationStep] = useState(false); // To track if we're in the verification step
    const [verificationCode, setVerificationCode] = useState(''); // Holds the user's input for the verification code
    const [statusMessage, setStatusMessage] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    // Handle input change
    const handleChange = (e) => {
        const { name, value, type } = e.target;
        setFormData({ ...formData, [name]: type === 'radio' ? value : value });
        console.log("course selected: ",formData.course);
    };

    // Handle verification code input change
    const handleVerificationChange = (e) => {
        setVerificationCode(e.target.value);
    };

    // Handle form submission
    const handleSubmit = async (e) => {
        e.preventDefault();

        // Validate required fields
        if (!formData.phone || !formData.pin || !formData.firstName || !formData.lastName || !formData.role) {
            setStatusMessage('Please fill out all required fields.');
            return;
        }

        try {
            setIsSubmitting(true);
            setStatusMessage('');

            // Send form data to the registration API
            const response = await axios.post('https://api.acework.ai/register', formData);

            if (response.status === 200) {
                // Proceed to the verification step
                setIsVerificationStep(true);
                setStatusMessage('Registration successful! A verification code has been sent to your phone.');
            }
        } catch (error) {
            setStatusMessage('An error occurred while registering. Please try again later.');
        } finally {
            setIsSubmitting(false);
        }
    };

    // Handle verification submission
    const handleVerify = async (e) => {
        e.preventDefault();

        if (!verificationCode) {
            setStatusMessage('Please enter the verification code.');
            return;
        }

        try {
            setIsSubmitting(true);
            setStatusMessage('');

            // Send verification code to the verification API
            const response = await axios.post('https://api.acework.ai/verify', {
                phone: formData.phone,
                code: verificationCode,
            });

            if (response.status === 200) {
                setStatusMessage('Verification successful! Your account is now fully registered.');
                setFormData({
                    phone: '',
                    pin: '',
                    firstName: '',
                    lastName: '',
                    role: '',
                    specialty: '',
                    nonMedText: '',
                    organize: '',
                    position: '',
                    dob: '',
                    course: '',
                    lineID: '',
                    facebook: '',
                    instagram: '',
                    email: '',
                    foodAllergy: '',
                    religion: '',
                    otherReligion: '',
                    note: '',
                });
                setIsVerificationStep(false);
                setVerificationCode('');
            }
        } catch (error) {
            setStatusMessage('Verification failed. Please try again.');
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <div className="container mt-5">
            <h2 className="text-center mb-4">Register สมัครสมาชิก</h2>

            {!isVerificationStep ? (
                <form onSubmit={handleSubmit}>
                    {/* Phone */}
                    <div className="mb-3">
                        <label htmlFor="phone" className="form-label">ID (เบอร์โทรศัพท์)</label>
                        <input
                            type="text"
                            id="phone"
                            name="phone"
                            maxLength="10"
                            className="form-control"
                            placeholder="Enter your phone number"
                            value={formData.phone}
                            onChange={handleChange}
                            required
                        />
                    </div>

                    {/* Pin */}
                    <div className="mb-3">
                        <label htmlFor="pin" className="form-label">Pin (6 Digit)</label>
                        <input
                            type="text"
                            id="pin"
                            name="pin"
                            maxLength="6"
                            className="form-control"
                            placeholder="Enter your PIN"
                            value={formData.pin}
                            onChange={handleChange}
                            required
                        />
                    </div>

                    {/* Name */}
                    <div className="row">
                        <div className="col-md-6 mb-3">
                            <label htmlFor="firstName" className="form-label">ชื่อ</label>
                            <input
                                type="text"
                                id="firstName"
                                name="firstName"
                                maxLength="50"
                                className="form-control"
                                placeholder="Enter your first name"
                                value={formData.firstName}
                                onChange={handleChange}
                                required
                            />
                        </div>
                        <div className="col-md-6 mb-3">
                            <label htmlFor="lastName" className="form-label">นามสกุล</label>
                            <input
                                type="text"
                                id="lastName"
                                name="lastName"
                                maxLength="50"
                                className="form-control"
                                placeholder="Enter your last name"
                                value={formData.lastName}
                                onChange={handleChange}
                                required
                            />
                        </div>
                    </div>

                    {/* Role */}
                    <div className="mb-3">
                        <label className="form-label">Role</label>
                        <div>
                            <div className="form-check">
                                <input
                                    type="radio"
                                    id="med"
                                    name="role"
                                    value="Med"
                                    className="form-check-input"
                                    onChange={handleChange}
                                    checked={formData.role === 'Med'}
                                />
                                <label htmlFor="med" className="form-check-label">Med</label>
                            </div>
                            <div className="form-check">
                                <input
                                    type="radio"
                                    id="nonMed"
                                    name="role"
                                    value="Non-Med"
                                    className="form-check-input"
                                    onChange={handleChange}
                                    checked={formData.role === 'Non-Med'}
                                />
                                <label htmlFor="nonMed" className="form-check-label">Non-Med</label>
                            </div>
                        </div>
                    </div>
                    {/* organization and position*/}
                    <div className="row">
                        <div className="col-md-6 mb-3">
                            <label className={"form-label"} htmlFor={"organization"}>หน่วยงาน</label>
                            <input
                                type="text"
                                id="organize"
                                name="organize"
                                className="form-control"
                                placeholder="Enter your organization"
                                value={formData.organize}
                                onChange={handleChange}
                            />
                        </div>
                        <div className="col-md-6 mb-3">
                            <label className={"form-label"} htmlFor={"position"}>ตำแหน่ง</label>
                            <input
                                type="text"
                                className="form-control"
                                id="position"
                                name="position"
                                placeholder="Enter your position"
                                value={formData.position}
                                onChange={handleChange}
                            />
                        </div>
                    </div>
                    <div className={"mb-3"}>
                        <label className="form-label" htmlFor="course">หลักสูตรที่เคยร่วม</label>
                        <input
                            type="text"
                            className="form-control"
                            id="course"
                            name="course"
                            value={formData.course}
                            onChange={handleChange}
                        />
                    </div>
                    {/*Line */}
                    <div className="mb-3">
                        <label htmlFor="line" className="lineAt">@Line</label>
                        <input
                            type="text"
                            maxLength="30"
                            className="form-control"
                            id="line"
                            name="line"
                            value={formData.line}
                            onChange={handleChange}/>
                    </div>
                    {/*instagram */}
                    <div className="mb-3">
                        <label htmlFor="instragram" className="form-label">Instagram</label>
                        <input
                            type="text"
                            className="form-control"
                            id="instagram"
                            name="instagram"
                            value={formData.instagram}
                            onChange={handleChange}
                        />
                    </div>
                    {/* */}
                    <div className="mb-3">
                        <label className="form-label">หลักสูตรที่เคยร่วม</label>
                        <select
                            id="course"
                            name="course"
                            className="form-select"
                            value={formData.course}
                            onChange={handleChange}
                            aria-label="your course select">
                            <option value="">-- Select a Course --</option>
                            {courses.map((course, index) => (
                                <option key={index} value={course.value}>
                                    {course.title}
                                </option>
                            ))}
                            {/*<CourseList courses={courses} />*/}
                        </select>

                    </div>


                    <button
                        type="submit"
                        className="btn btn-primary w-100"
                        disabled={isSubmitting}
                    >
                        {isSubmitting ? 'Registering...' : 'Register'}
                    </button>
                    <div className="row mb-3"></div>
                </form>
            ) : (
                <form onSubmit={handleVerify}>
                    <div className="mb-3">
                        <label htmlFor="verificationCode" className="form-label">Verification Code</label>
                        <input
                            type="text"
                            id="verificationCode"
                            name="verificationCode"
                            className="form-control"
                            placeholder="Enter the verification code sent to your phone"
                            value={verificationCode}
                            onChange={handleVerificationChange}
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        className="btn btn-primary w-100"
                        disabled={isSubmitting}
                    >
                        {isSubmitting ? 'Verifying...' : 'Verify'}
                    </button>
                </form>
            )}

            {statusMessage && (
                <div className={`mt-3 alert ${statusMessage.includes('successful') ? 'alert-success' : 'alert-danger'}`}>
                    {statusMessage}
                </div>
            )}
        </div>
    );
};

export default Register;