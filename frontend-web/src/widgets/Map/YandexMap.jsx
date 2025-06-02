import React, { useEffect, useState } from "react";
import loadReactifiedYmaps from "./ymaps";

const YandexMap = () => {
  const [ymapsReady, setYmapsReady] = useState(false);
  const [YMapComponents, setYMapComponents] = useState(null);

  useEffect(() => {
    const init = async () => {
      try {
        const ymaps = await loadReactifiedYmaps();
        setYMapComponents(ymaps);
        setYmapsReady(true);
      } catch (err) {
        console.error("Ошибка загрузки Yandex Maps:", err);
      }
    };

    init();
  }, []);

  if (!ymapsReady || !YMapComponents) return <div>Загрузка карты...</div>;

  const { YMap, YMapDefaultSchemeLayer, YMapDefaultFeaturesLayer, YMapMarker } =
    YMapComponents;

  return (
    <YMap
      location={{ center: [37.588144, 55.733842], zoom: 10 }}
      mode="vector"
      style={{ width: "100%", height: "500px" }}
    >
      <YMapDefaultSchemeLayer />
      <YMapDefaultFeaturesLayer />
      <YMapMarker coordinates={[37.588144, 55.733842]}>
        <div style={{ backgroundColor: "white", padding: 4 }}>Москва</div>
      </YMapMarker>
    </YMap>
  );
};

export default YandexMap;
