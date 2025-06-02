import React from "react";
import ReactDOM from "react-dom/client";

const loadYmaps3 = () => {
  return new Promise((resolve, reject) => {
    if (window.ymaps3) {
      resolve(window.ymaps3);
      return;
    }

    const script = document.createElement("script");
    script.src = `https://api-maps.yandex.ru/v3/?apikey=${
      import.meta.env.VITE_YANDEX_MAP_API_KEY
    }&lang=ru_RU`;
    script.type = "text/javascript";
    script.onload = () => {
      if (window.ymaps3) {
        resolve(window.ymaps3);
      } else {
        reject(new Error("ymaps3 failed to load"));
      }
    };
    script.onerror = reject;
    document.head.appendChild(script);
  });
};

const loadReactifiedYmaps = async () => {
  const ymaps3 = await loadYmaps3();
  const [ymaps3React] = await Promise.all([
    ymaps3.import("@yandex/ymaps3-reactify"),
    ymaps3.ready,
  ]);

  const reactify = ymaps3React.reactify.bindTo(React, ReactDOM);
  const { YMap, YMapDefaultSchemeLayer, YMapDefaultFeaturesLayer, YMapMarker } =
    reactify.module(ymaps3);

  return {
    YMap,
    YMapDefaultSchemeLayer,
    YMapDefaultFeaturesLayer,
    YMapMarker,
    reactify,
  };
};

export default loadReactifiedYmaps;
