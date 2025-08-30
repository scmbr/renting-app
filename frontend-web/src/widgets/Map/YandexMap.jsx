import React, { useEffect, useRef } from "react";
import { YMaps, Map, Placemark, Clusterer } from "@pbe/react-yandex-maps";
import { useCityStore } from "@/stores/useCityStore";
import { useMapStore } from "@/stores/useMapStore";
import { getCoordsByCity } from "@/shared/constants/cities";
import { useFiltersStore } from "@/stores/useFiltersStore";
import { useNavigate } from "react-router-dom";
import Spinner from "@/widgets/Spinner/Spinner";

export const YandexMap = ({ markerPosition, onSelect, adverts = [] }) => {
  const mapRef = useRef(null);
  const city = useCityStore((state) => state.city);
  const updateFilter = useFiltersStore((state) => state.updateFilter);
  const navigate = useNavigate();

  const mapState = useMapStore((state) => state.mapState);
  const setMapState = useMapStore((state) => state.setMapState);

  useEffect(() => {
    const loadCoords = async () => {
      if (mapState?.city === city) return;

      const coords = await getCoordsByCity(city);
      if (coords) setMapState(city, coords, 12); // default zoom
    };
    loadCoords();
  }, [city, mapState, setMapState]);

  const handleMapClick = (e) => {
    const coords = e.get("coords");
    onSelect?.(coords);
    updateFilter("lat", undefined);
    updateFilter("lng", undefined);

    // Обновляем центр карты в сторе
    setMapState(city, coords, mapState.zoom);
  };

  const handleBoundsChange = (event) => {
    const newCenter = event.get("newCenter");
    const newZoom = event.get("newZoom");

    setMapState(city, newCenter, newZoom);
  };

  const handleMarkerClick = (advert) => {
    if (advert?.id != null) {
      navigate(`/advert/${advert.id}`);
    }
  };

  if (!mapState?.coords) {
    return <Spinner />;
  }

  const mapCenter = mapState.coords;
  const mapZoom = mapState.zoom ?? 12;

  return (
    <YMaps>
      <Map
        state={{ center: mapCenter, zoom: mapZoom }}
        width="100%"
        height="100%"
        onClick={handleMapClick}
        onBoundsChange={handleBoundsChange}
        instanceRef={mapRef}
      >
        <Clusterer
          options={{
            preset: "islands#invertedBlackClusterIcons",
            groupByCoordinates: false,
            clusterDisableClickZoom: false,
            clusterOpenBalloonOnClick: false,
          }}
          onClick={(e) => {
            const cluster = e.get("target");
            const geoObjects = cluster.getGeoObjects();

            if (geoObjects.length === 0) return;

            const firstCoords = geoObjects[0].geometry.getCoordinates();
            const allSame = geoObjects.every((obj) => {
              const coords = obj.geometry.getCoordinates();
              return (
                coords[0] === firstCoords[0] && coords[1] === firstCoords[1]
              );
            });

            if (allSame) {
              updateFilter("lat", firstCoords[0]);
              updateFilter("lng", firstCoords[1]);
            } else {
              updateFilter("lat", undefined);
              updateFilter("lng", undefined);
            }
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
                properties={{ iconCaption: `${advert.rent} ₽` }}
                options={{ preset: "islands#dotIcon", iconColor: "#8b51ff" }}
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
