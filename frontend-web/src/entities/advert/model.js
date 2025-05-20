import api from '@/shared/api/axios';

export async function fetchAdverts(filters) {
  const params = {};
  if (filters.city) params.city = filters.city;
  const response = await api.get('/adverts', { params });
  return response.data;
}