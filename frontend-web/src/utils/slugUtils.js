import { transliterate, slugify as transliterateToSlug } from "transliteration";


const slugToCityMap = {
  moskva: "Москва",
  kazan: "Казань",
  "sankt-peterburg": "Санкт-Петербург",
};

export const slugify = (city) => transliterateToSlug(city).toLowerCase();
export const getCityFromSlug = (slug) => {
  return slugToCityMap[slug] || transliterate(slug); // fallback — пусть будет хоть что-то
};
