import React, { useEffect, useState } from "react";
import { YMaps, Map, Placemark } from "@pbe/react-yandex-maps";

const DEFAULT_COORDS = [55.75396, 37.620393]; // Москва, как в 2GIS

export const MapGLForm = ({ center, markerPosition }) => {
  const [mapCenter, setMapCenter] = useState(center || DEFAULT_COORDS);
  const [zoom, setZoom] = useState(17);

  useEffect(() => {
    if (center) {
      setMapCenter(center);
    }
  }, [center?.[0], center?.[1]]);

  return (
    <YMaps>
      <Map
        defaultState={{ center: mapCenter, zoom }}
        state={{ center: mapCenter, zoom }}
        width="100%"
        height="300px"
      >
        {markerPosition && (
          <Placemark
            geometry={markerPosition}
            options={{
              preset: "islands#dotIcon",
              iconColor: "#8b51ff",
            }}
          />
        )}
      </Map>
    </YMaps>
  );
};
