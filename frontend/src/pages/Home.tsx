import React from "react"
import RegistrationForm from "../components/RegistrationForm"
import RegistrationList from "../components/RegistrationList"

const Home: React.FC = () => {
  return (
    <div>
      <RegistrationForm />
      <RegistrationList />
    </div>
  )
}

export default Home
