import React, { useEffect, useRef, useState } from "react";
import { YMaps, Map, Placemark, Clusterer } from "@pbe/react-yandex-maps";
import { useCityStore } from "@/stores/useCityStore";
import { getCoordsByCity } from "@/shared/constants/cities";
import { useFiltersStore } from "@/stores/useFiltersStore";
import { useNavigate } from "react-router-dom";
import Spinner from "@/widgets/Spinner/Spinner";

export const YandexMap = ({ markerPosition, onSelect, adverts = [] }) => {
  const mapRef = useRef(null);
  const city = useCityStore((state) => state.city);
  const updateFilter = useFiltersStore((state) => state.updateFilter);
  const navigate = useNavigate();

  const [mapCenter, setMapCenter] = useState(null);
  const [zoom, setZoom] = useState(12);

  useEffect(() => {
    const loadCoords = async () => {
      const coords = await getCoordsByCity(city);
      if (coords) {
        setMapCenter(coords);
      }
    };
    loadCoords();
  }, [city]);

  const handleMapClick = (e) => {
    const coords = e.get("coords");
    onSelect?.(coords);
    updateFilter("lat", coords[0]);
    updateFilter("lng", coords[1]);
  };

  const handleMarkerClick = (advert) => {
    if (advert?.id != null) {
      navigate(`/advert/${advert.id}`);
    }
  };

  if (!mapCenter) {
    return <Spinner />;
  }

  return (
    <YMaps>
      <Map
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
