export const slugToName = (slug) => {
  const cityEntry = Object.entries(cityMappings).find(([_, s]) => s === slug);
  return (
    cityEntry?.[0] ||
    slug.replace(/-/g, " ").replace(/(?:^|\s)\S/g, (a) => a.toUpperCase())
  );
};

export const nameToSlug = (cityName) => {
  const slugFromMap = cityMappings[cityName];
  if (slugFromMap) return slugFromMap;

  return cityName
    .toLowerCase()
    .normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-+|-+$/g, "");
};
export async function getCoordsByCity(cityName) {
  const apiKey = import.meta.env.VITE_YANDEX_GEOCODER_KEY;
  const baseUrl = "https://geocode-maps.yandex.ru/1.x/";
  const url = `${baseUrl}?format=json&apikey=${apiKey}&geocode=${encodeURIComponent(
    cityName
  )}`;

  try {
    const response = await fetch(url);
    const data = await response.json();

    const geoObject =
      data?.response?.GeoObjectCollection?.featureMember?.[0]?.GeoObject;

    if (!geoObject || !geoObject.Point?.pos) {
      return null;
    }
    const [lon, lat] = geoObject.Point.pos.split(" ").map(Number);
    return [lon, lat];
  } catch (err) {
    return null;
  }
}
export const cityMappings = {
  Москва: "moskva",
  "Санкт-Петербург": "spb",
  Новосибирск: "novosibirsk",
  Екатеринбург: "ekaterinburg",
  Казань: "kazan",
  "Нижний Новгород": "nnov",
  Челябинск: "chelyabinsk",
  Самара: "samara",
  Омск: "omsk",
  "Ростов-на-Дону": "rostov",
  Уфа: "ufa",
  Красноярск: "krasnoyarsk",
  Пермь: "perm",
  Воронеж: "voronezh",
  Волгоград: "volgograd",
  Краснодар: "krasnodar",
  Саратов: "saratov",
  Тюмень: "tyumen",
  Тольятти: "tolyatti",
  Ижевск: "izhevsk",
  Барнаул: "barnaul",
  Ульяновск: "ulyanovsk",
  Иркутск: "irkutsk",
  Хабаровск: "khabarovsk",
  Ярославль: "yaroslavl",
  Владивосток: "vladivostok",
  Махачкала: "makhachkala",
  Томск: "tomsk",
  Оренбург: "orenburg",
  Кемерово: "kemerovo",
  Новокузнецк: "novokuznetsk",
  Рязань: "ryazan",
  Астрахань: "astrakhan",
  "Набережные Челны": "nabchelny",
  Пенза: "penza",
  Липецк: "lipetsk",
  Киров: "kirov",
  Чебоксары: "cheboksary",
  Тула: "tula",
  Калининград: "kaliningrad",
  Балашиха: "balashikha",
  Курск: "kursk",
  Севастополь: "sevastopol",
  Сочи: "sochi",
  Ставрополь: "stavropol",
  "Улан-Удэ": "ulanude",
  Магнитогорск: "magnitogorsk",
  Тверь: "tver",
  Иваново: "ivanovo",
  Брянск: "bryansk",
  Архангельск: "arkhangelsk",
  Владимир: "vladimir",
  Чита: "chita",
  Симферополь: "simferopol",
  Грозный: "grozny",
  Курган: "kurgan",
  Орёл: "orel",
  Волжский: "volzhsky",
  Смоленск: "smolensk",
  Мурманск: "murmansk",
  Владикавказ: "vladikavkaz",
  Якутск: "yakutsk",
  Саранск: "saransk",
  Череповец: "cherepovets",
  Вологда: "vologda",
  Орск: "orsk",
  Стерлитамак: "sterlitamak",
  Глазов: "glazov",
  "Новый Уренгой": "newurengoy",
  Абакан: "abakan",
  Нальчик: "nalchik",
  Находка: "nakhodka",
  "Йошкар-Ола": "yoshkarola",
  Бийск: "biysk",
  Рубцовск: "rubtsovsk",
  Благовещенск: "blagoveshchensk",
  Прокопьевск: "prokopyevsk",
  "Старый Оскол": "staryoskol",
  Златоуст: "zlatoust",
  Миасс: "miass",
  "Ленинск-Кузнецкий": "leninsk-kuznetsky",
  Сыктывкар: "syktyvkar",
  Канск: "kansk",
  Новороссийск: "novorossiysk",
  Шахты: "shakhty",
  Нижневартовск: "nizhnevartovsk",
  Дзержинск: "dzierzynsk",
  Октябрьский: "oktyabrsky",
  Элиста: "elista",
  Армавир: "armavir",
  Бердск: "berdsk",
  Назрань: "nazran",
  Ангарск: "angarsk",
  Уссурийск: "ussuriysk",
  Королёв: "korolyov",
  Петрозаводск: "petrozavodsk",
  Сызрань: "syzran",
  Норильск: "norilsk",
  Зеленодольск: "zelenodolsk",
  Междуреченск: "mezhdurechensk",
  Альметьевск: "almetyevsk",
  Копейск: "kopeysk",
  Майкоп: "maykop",
  Балаково: "balakovo",
  Кызыл: "kyzyl",
  Железногорск: "zheleznogorsk",
  Северодвинск: "severodvinsk",
  Артём: "artem",
  Новочебоксарск: "novocheboksarsk",
  Серпухов: "serpukhov",
  Димитровград: "dimitrovgrad",
  "Каменск-Уральский": "kamensk-uralsky",
  Нефтеюганск: "nefteyugansk",
  Первоуральск: "pervouralsk",
  "Орехово-Зуево": "orekhovo-zuevo",
  Нефтекамск: "neftekamsk",
  Дербент: "derbent",
  Черкесск: "cherkessk",
  Озёрск: "ozersk",
  Бугульма: "bugulma",
  Новошахтинск: "novoshakhtinsk",
  Евпатория: "evpatoria",
  Кисловодск: "kislovodsk",
  Долгопрудный: "dolgoprudny",
  Жуковский: "zhukovsky",
  Реутов: "reutov",
  Пушкино: "pushkino",
  Раменское: "ramenskoye",
  Обнинск: "obninsk",
  Домодедово: "domodedovo",
  "Сергиев Посад": "sergiev-posad",
  Электросталь: "electrostal",
  Арзамас: "arzamas",
  Клин: "klin",
  "Наро-Фоминск": "naro-fominsk",
  Щёлково: "shchelkovo",
  Фрязино: "fryazino",
  Лобня: "lobnya",
  Дубна: "dubna",
  "Павловский Посад": "pavlovsky-posad",
  Коломна: "kolomna",
  Муром: "murom",
  Егорьевск: "egoryevsk",
  Воскресенск: "voskresensk",
  Шатура: "shatura",
  Лыткарино: "lytkarino",
  Черноголовка: "chernogolovka",
  Дмитров: "dmitrov",
  Солнечногорск: "solnechnogorsk",
  Ивантеевка: "ivanteevka",
  Краснознаменск: "krasnoznamensk",
  Видное: "vidnoe",
  Феодосия: "feodosiya",
  Керчь: "kerch",
};
