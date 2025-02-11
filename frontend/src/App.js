import React,{useState} from 'react';
import 'bootstrap/dist/js/bootstrap.min.js'
import Register from "./components/Register";
import News from "./components/News";
import Attend from "./components/Attend";
import Calendar from "./components/Calendar";
import Contact from "./components/Contact";
import Home from "./components/Home";
import EventManagement from "./components/EventManagement";
function App() {
    const [currentPage, setCurrentPage] = useState('home'); // State for tracking the current page

    // Function to render the selected component
    const renderPage = () => {
        switch (currentPage) {
            case 'register':
                return <Register />;
            case 'news':
                return <News />;
            case 'attend':
                return <Attend />;
            case 'calendar':
                return <Calendar />;
            case 'contact':
                return <Contact />;
            case 'home':
                return <Home />;
            case 'eventManagement':
                return <EventManagement />;
            default:
                return <div className="container mt-5"><h1>404 - Page Not Found</h1></div>
        }
    }


  return (
    <div className="App">
        {/* Navbar */}
        <nav className="navbar navbar-expand-lg navbar-dark bg-primary">
            <div className="container">
                <a className="navbar-brand" href="#">Doctor App</a>
                <button
                    className="navbar-toggler"
                    type="button"
                    data-bs-toggle="collapse"
                    data-bs-target="#navbarNav"
                    aria-controls="navbarNav"
                    aria-expanded="false"
                    aria-label="Toggle navigation"
                >
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarNav">
                    <ul className="navbar-nav ms-auto">
                        <li className="nav-item">
                            {/*<a className="nav-link" href="#register">Register สมัครสมาชิก</a>*/}
                            <a
                                className={`nav-link ${currentPage === 'home' ? 'active' : ''}`}
                                href="#"
                                onClick={() => setCurrentPage('home')}
                            >
                                Home
                            </a>
                        </li>
                        {/*event management menu*/}
                        <li className="nav-item">
                            {/*<a className="nav-link" href="#register">Register สมัครสมาชิก</a>*/}
                            <a
                                className={`nav-link ${currentPage === 'eventManagement' ? 'active' : ''}`}
                                href="#"
                                onClick={() => setCurrentPage('eventManagement')}
                            >
                                Event Management
                            </a>
                        </li>


                        <li className="nav-item">
                            <a
                                className={`nav-link ${currentPage === 'register' ? 'active' : ''}`}
                                href="#"
                                onClick={() => setCurrentPage('register')}
                            >
                                Register
                            </a>
                        </li>
                        <li className="nav-item">
                            <a className={`nav-link ${currentPage === 'news' ? 'active' : ''}`} href="#" onClick={() => setCurrentPage('news')}
                        >
                                News
                            </a>
                        </li>
                        <li className="nav-item">
                        {/*<a className="nav-link" href="#attend">Attend ร่วมงานแพทยอาสา</a>*/}
                            <a
                                className={`nav-link ${currentPage === 'attend' ? 'active' : ''}`}
                                href="#"
                                onClick={() => setCurrentPage('attend')}
                            >
                                Attend
                            </a>
                        </li>
                        <li className="nav-item">
                            {/*<a className="nav-link" href="#calendar">Calendar ปฏิทินงาน</a>*/}
                            <a className={`nav-link ${currentPage === 'calendar' ? 'active' : ''}`} href="#" onClick={() => setCurrentPage('calendar')} >
                                Calendar
                            </a>
                        </li>
                        <li className="nav-item">
                            {/*<a className="nav-link" href="#contact">Contact Us ติดต่อ</a>*/}
                            <a className={`nav-link ${currentPage === 'contact' ? 'active' : ''}`} href="#" onClick={() => setCurrentPage('contact')} >
                                Contact Us
                            </a>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
        {/* Main Content */}
        <div className="container mt-5">
            {renderPage()}

            {/*
            <section id="register" className="mb-5">

                <h2>Register สมัครสมาชิก</h2>
                <p>Here you can register for our services.</p>
            </section>
            <section id="news" className="mb-5">
                <h2>News ประชาสัมพันธ์</h2>
                <p>Latest updates and announcements.</p>
            </section>
            <section id="attend" className="mb-5">
                <h2>Attend ร่วมงานแพทยอาสา</h2>
                <p>Join our medical volunteer program.</p>
            </section>
            <section id="calendar" className="mb-5">
                <h2>Calendar ปฏิทินงาน</h2>
                <p>View upcoming events and schedules.</p>
            </section>
            <section id="contact" className="mb-5">
                <h2>Contact Us ติดต่อ</h2>
                <p>Get in touch with us for more information.</p>
            </section>
            */}
        </div>

        {/* Footer */}
        <footer className="bg-dark text-white text-center py-3 mt-auto">
            <div className="container">
                <p className="mb-0">&copy; 2025 My App. All rights reserved.</p>
            </div>
        </footer>

    </div>

  );
}

export default App;
