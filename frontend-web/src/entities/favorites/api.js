import api from "@/shared/api/axios";

export async function fetchFavorites() {
  const response = await api.get("/favorites");
  return response.data;
}
export async function addToFavorites(advert_id) {
  return api.post(`/favorites`, { advert_id });
}

export async function removeFromFavorites(id) {
  return api.delete(`/favorites/${id}`);
}
