import { slugToName } from "@/shared/constants/cities";
import { useCityStore } from "@/stores/useCityStore";

export function initCityFromUrl() {
  const params = new URLSearchParams(window.location.search);
  const slug = params.get("city");

  const city = slug ? slugToName(slug) : "Москва";

  const setCity = useCityStore.getState().setCity;
  setCity(city);
}
