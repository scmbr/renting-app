import React, { useEffect, useRef } from "react";
import { load } from "@2gis/mapgl";

const DEFAULT_COORDS = [37.620393, 55.75396];

export const MapGLForm = ({ center, markerPosition }) => {
  const containerRef = useRef(null);
  const mapRef = useRef(null);
  const markerRef = useRef(null);

  useEffect(() => {
    let destroyed = false;

    const initMap = async () => {
      const mapglAPI = await load();
      if (destroyed || !containerRef.current) return;

      const map = new mapglAPI.Map(containerRef.current, {
        key: import.meta.env.VITE_2GIS_MAP_API_KEY,
        center: center || DEFAULT_COORDS,
        zoom: 14,
      });
      mapRef.current = map;

      if (markerPosition) {
        const marker = new mapglAPI.Marker(map, {
          coordinates: markerPosition,
        });
        markerRef.current = marker;
      }
    };

    initMap();

    return () => {
      destroyed = true;
      markerRef.current?.destroy();
      mapRef.current?.destroy();
    };
  }, []);

  useEffect(() => {
    if (mapRef.current && center) {
      mapRef.current.setCenter(center);
    }

    if (markerRef.current) {
      if (markerPosition) {
        markerRef.current.setCoordinates(markerPosition);
      } else {
        markerRef.current.destroy();
        markerRef.current = null;
      }
    } else if (markerPosition && mapRef.current) {
      (async () => {
        const mapglAPI = await load();
        markerRef.current = new mapglAPI.Marker(mapRef.current, {
          coordinates: markerPosition,
        });
      })();
    }
  }, [center?.[0], center?.[1], markerPosition?.[0], markerPosition?.[1]]);

  return (
    <div
      ref={containerRef}
      style={{ width: "100%", height: "300px", borderRadius: "8px" }}
    />
  );
};
