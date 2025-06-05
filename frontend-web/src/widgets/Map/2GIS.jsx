import { load } from "@2gis/mapgl";
import React, { useEffect, useRef, useState } from "react";
import { useCityStore } from "@/stores/useCityStore";
import { getCoordsByCity } from "@/shared/constants/cities";
import { Clusterer } from "@2gis/mapgl-clusterer";

export const MapGL = ({ markerPosition, onSelect, adverts = [] }) => {
  const containerRef = useRef(null);
  const mapRef = useRef(null);
  const mapglAPIRef = useRef(null);
  const clustererRef = useRef(null);
  const city = useCityStore((state) => state.city);
  const [isMapReady, setIsMapReady] = useState(false);
  console.log(adverts);
  useEffect(() => {
    if (clustererRef.current) {
      clustererRef.current.destroy();
      clustererRef.current = null;
    }
  }, [city]);

  useEffect(() => {
    let destroyed = false;

    const initMap = async () => {
      const mapglAPI = await load();
      mapglAPIRef.current = mapglAPI;

      let coords = await getCoordsByCity(city);
      if (!coords) {
        coords = [60.751244, 59.618423];
      }

      if (destroyed || !containerRef.current) return;

      const map = new mapglAPI.Map(containerRef.current, {
        center: coords,
        zoom: 13,
        key: import.meta.env.VITE_2GIS_MAP_API_KEY,
        style: "96072560-ad86-46df-8126-b6f478fc6010",
      });

      mapRef.current = map;

      setTimeout(() => {
        map.resize();
      }, 0);

      map.on("click", (event) => {
        const coords = event.lngLat;
        onSelect?.(coords);
      });

      setIsMapReady(true);
    };

    initMap();

    return () => {
      destroyed = true;
      clustererRef.current?.destroy();
      clustererRef.current = null;
      mapRef.current?.destroy();
      mapRef.current = null;
      mapglAPIRef.current = null;
    };
  }, [city, onSelect]);

  useEffect(() => {
    if (
      !isMapReady ||
      !mapRef.current ||
      !mapglAPIRef.current ||
      !adverts?.length
    )
      return;

    if (clustererRef.current) {
      clustererRef.current.destroy();
      clustererRef.current = null;
    }

    const clusterer = new Clusterer(mapRef.current, {
      clusterStyle: {
        icon: "/icons/cluster.png",
        labelColor: "#ffffff",
        size: [30, 30],
        labelFontSize: 12,
        offset: [0, -20],
      },
      radius: 30,
    });
    clustererRef.current = clusterer;

    const markers = adverts
      .map((advert) => {
        const apt = advert.apartment;
        if (
          apt &&
          typeof apt.latitude === "number" &&
          typeof apt.longitude === "number"
        ) {
          return {
            coordinates: [apt.longitude, apt.latitude],
            icon: "/icons/apartment.png",
            size: [45, 24],
            label: {
              text: advert.rent + " ₽",
              color: "#fff",
              fontSize: 11,
            },
            data: advert,
          };
        }
        return null;
      })
      .filter(Boolean);

    clusterer.load(markers);

    clusterer.on("click", (event) => {
      if (event.target.type === "cluster") {
        console.log("Кластер кликнут", event.target.data);
      } else if (event.target.type === "marker") {
        console.log("Маркер кликнут", event.target.data);
      }
    });
  }, [adverts, isMapReady, city]);

  useEffect(() => {
    if (!mapRef.current || !city) return;

    getCoordsByCity(city).then((coords) => {
      if (coords) {
        mapRef.current.setCenter(coords);
      }
    });
  }, [city]);

  return (
    <div
      ref={containerRef}
      style={{ width: "100%", height: "100%", borderRadius: "8px" }}
    />
  );
};
