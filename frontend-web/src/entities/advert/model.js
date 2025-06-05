import api from "@/shared/api/axios";

export async function fetchAdverts(filters) {
  const params = {};
  for (const [key, value] of Object.entries(filters)) {
    if (value !== undefined && value !== "") {
      params[key] = value;
    }
  }
  const response = await api.get("/adverts", { params });
  return response.data;
}

export async function fetchAdvertById(id) {
  const response = await api.get(`/adverts/${id}`);
  return response.data;
}
