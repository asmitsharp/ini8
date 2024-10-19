import React from "react"
import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import Navbar from "./components/Navbar"
import RegistrationList from "./components/RegistrationList"
import RegistrationForm from "./components/RegistrationForm"
import { ToastContainer } from "react-toastify"
import "react-toastify/dist/ReactToastify.css"

const App: React.FC = () => {
  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <Navbar />
        <div className="container mx-auto px-4 py-8">
          <Routes>
            <Route path="/" element={<RegistrationList />} />
            <Route path="/new" element={<RegistrationForm />} />
            <Route path="/edit/:id" element={<RegistrationForm />} />
          </Routes>
        </div>
        <ToastContainer />
      </div>
    </Router>
  )
}

export default App
