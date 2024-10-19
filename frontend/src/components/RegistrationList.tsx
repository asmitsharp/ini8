import React, { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { toast } from "react-toastify"
import { fetchRegistrations, deleteRegistration } from "../services/api"
import { formatDate } from "../utils/dateUtils"
import { Registration } from "../types"

const RegistrationList: React.FC = () => {
  const [registrations, setRegistrations] = useState<Registration[]>([])
  const [loading, setLoading] = useState<boolean>(true)
  const navigate = useNavigate()

  useEffect(() => {
    loadRegistrations()
  }, [])

  const loadRegistrations = async (): Promise<void> => {
    try {
      const data = await fetchRegistrations()
      setRegistrations(data)
      setLoading(false)
    } catch (error) {
      toast.error("Failed to load registrations")
      setLoading(false)
    }
  }

  const handleDelete = async (id: number): Promise<void> => {
    if (window.confirm("Are you sure you want to delete this registration?")) {
      try {
        await deleteRegistration(id)
        toast.success("Registration deleted successfully")
        loadRegistrations()
      } catch (error) {
        toast.error("Failed to delete registration")
      }
    }
  }

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
      </div>
    )
  }

  return (
    <div className="bg-white rounded-lg shadow overflow-hidden">
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Name
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Email
            </th>
            <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Date of Birth
            </th>
            <th className="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
              Actions
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {registrations.map((registration) => (
            <tr key={registration.id}>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                {registration.name}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {registration.email}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {formatDate(registration.dob)}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button
                  onClick={() => navigate(`/edit/${registration.id}`)}
                  className="text-blue-600 hover:text-blue-900 mr-4"
                >
                  Edit
                </button>
                <button
                  onClick={() => handleDelete(registration.id)}
                  className="text-red-600 hover:text-red-900"
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default RegistrationList
