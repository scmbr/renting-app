import { load } from "@2gis/mapgl";
import React, { useEffect, useRef } from "react";

export const MapGL = () => {
  const containerRef = useRef(null);

  useEffect(() => {
    let map;

    load().then((mapglAPI) => {
      if (containerRef.current && !containerRef.current.hasChildNodes()) {
        map = new mapglAPI.Map(containerRef.current, {
          center: [60.751244, 59.618423],
          zoom: 13,
          key: import.meta.env.VITE_2GIS_MAP_API_KEY,
        });
      }
    });

    return () => map?.destroy();
  }, []);

  return <div ref={containerRef} style={{ width: "100%", height: "100%" }} />;
};
