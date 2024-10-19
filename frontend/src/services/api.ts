import { Registration, FormData } from "../types"

const API_BASE_URL =
  process.env.REACT_APP_API_BASE_URL || "http://localhost:3000/api"

export const fetchRegistrations = async (): Promise<Registration[]> => {
  const response = await fetch(`${API_BASE_URL}/registrations`)
  if (!response.ok) throw new Error("Failed to fetch registrations")
  return response.json()
}

export const getRegistration = async (id: number): Promise<Registration> => {
  const response = await fetch(`${API_BASE_URL}/registrations/${id}`)
  if (!response.ok) throw new Error("Failed to fetch registration")
  return response.json()
}

export const createRegistration = async (
  data: FormData
): Promise<Registration> => {
  const response = await fetch(`${API_BASE_URL}/registrations`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
  if (!response.ok) throw new Error("Failed to create registration")
  return response.json()
}

export const updateRegistration = async (
  id: number,
  data: FormData
): Promise<Registration> => {
  const response = await fetch(`${API_BASE_URL}/registrations/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
  if (!response.ok) throw new Error("Failed to update registration")
  return response.json()
}

export const deleteRegistration = async (id: number): Promise<void> => {
  const response = await fetch(`${API_BASE_URL}/registrations/${id}`, {
    method: "DELETE",
  })
  if (!response.ok) throw new Error("Failed to delete registration")
}
