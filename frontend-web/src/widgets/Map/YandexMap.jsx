import React, { useEffect, useRef, useState } from "react";
import { YMaps, Map, Placemark, Clusterer } from "@pbe/react-yandex-maps";
import { useCityStore } from "@/stores/useCityStore";
import { getCoordsByCity } from "@/shared/constants/cities";
import { useFiltersStore } from "@/stores/useFiltersStore";
import { useNavigate } from "react-router-dom";

export const YandexMap = ({ markerPosition, onSelect, adverts = [] }) => {
  const mapRef = useRef(null);
  const city = useCityStore((state) => state.city);
  const updateFilter = useFiltersStore((state) => state.updateFilter);
  const navigate = useNavigate();
  const [mapCenter, setMapCenter] = useState([55.751574, 37.573856]); 
  const [zoom, setZoom] = useState(12);

  useEffect(() => {
    const loadCoords = async () => {
      const coords = await getCoordsByCity(city);

      if (coords) {
        setMapCenter(coords);
      }
      console.log(coords[1]);
    };
    loadCoords();
  }, [city]);

  const handleMapClick = (e) => {
    const coords = e.get("coords");
    onSelect?.(coords);
    updateFilter("lat", undefined);
    updateFilter("lng", undefined);
    console.log("Клик по карте, координаты очищены");
  };

  const handleMarkerClick = (advert) => {
    if (advert?.id) {
      navigate(`/advert/${advert.id}`);
    }
  };

  return (
    <YMaps>
      <Map
        defaultState={{ center: mapCenter, zoom }}
        state={{ center: mapCenter, zoom }}
        width="100%"
        height="100%"
        onClick={handleMapClick}
        instanceRef={mapRef}
      >
        <Clusterer
          options={{
            preset: "islands#invertedBlackClusterIcons", 
            groupByCoordinates: false,
            clusterDisableClickZoom: false,
            clusterOpenBalloonOnClick: false,
            clusterColor: "#000000ff",
            
          }}
        >
        {adverts.map((advert) => {
          const apt = advert.apartment;
          if (
            !apt ||
            typeof apt.latitude !== "number" ||
            typeof apt.longitude !== "number"
          )
            return null;

          return (
            <Placemark
              key={advert.id}
              geometry={[apt.latitude, apt.longitude]}
              properties={{
                iconCaption: `${advert.rent} ₽`,
              }}
              options={{
                preset: "islands#dotIcon",
                iconColor: "#8b51ff",
              }}
              onClick={() => handleMarkerClick(advert)}
            />
            
          );
        })}
        </Clusterer>

        {markerPosition && (
          <Placemark
            geometry={markerPosition}
            options={{ preset: "islands#redDotIcon" }}
          />
        )}
      </Map>
    </YMaps>
  );
};
