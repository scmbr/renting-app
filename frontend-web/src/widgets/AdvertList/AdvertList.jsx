// AdvertList.jsx
import React, { useEffect, useState } from 'react';
import { fetchAdverts } from '@/entities/advert/model';
import AdvertCard from '@/entities/advert/components/AdvertCard';

const AdvertList = ({ filters }) => {
  const [adverts, setAdverts] = useState([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    setLoading(true);
    setError(null);

    fetchAdverts(filters)
      .then(data => {
        setAdverts(data.adverts);
        setTotal(data.total);
      })
      .catch(err => {
        setError('Ошибка при загрузке объявлений');
        console.error(err);
      })
      .finally(() => setLoading(false));
  }, [filters]);

  if (loading) return <p>Загрузка...</p>;
  if (error) return <p className="text-red-500">{error}</p>;

  return (
    <div>
      <h2>Найдено: {total}</h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {adverts.map(ad => (
          <AdvertCard key={ad.id} advert={ad} />
        ))}
      </div>
    </div>
  );
};

export default AdvertList;
