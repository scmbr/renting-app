import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { fetchAdvertById } from "@/entities/advert/model";
import styles from "./AdvertPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import Spinner from "@/widgets/Spinner/Spinner.jsx";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";
import OwnerModal from "@/widgets/OwnerModal/OwnerModal.jsx";
import api from "@/shared/api/axios";
import { MapGLForm } from "@/widgets/Map/MapGLForm";
import Footer from "@/widgets/Footer/Footer";
import clsx from "clsx";
import FeatureCardStatic from "@/entities/feature-card-static/FeatureCardStatic";

import AdvertCard from "@/entities/advert/components/AdvertCard";
const fetchSimilarAdverts = async (advert) => {
  const params = {
    city: advert.apartment.city,
    district: advert.apartment.district,
    rooms: advert.apartment.rooms,
    price_from: Math.floor(advert.rent * 0.8),
    price_to: Math.ceil(advert.rent * 1.2),
    limit: 10,
    offset: 0,
    sort_by: "created_at",
    order: "desc",
  };
  const response = await api.get("/adverts", { params });

  // Получаем массив из ответа
  const adverts = response.data.adverts || [];

  // Убираем текущее объявление из похожих
  return adverts.filter((a) => a.id !== advert.id);
};
const Stars = ({ rating }) => {
  const stars = [];
  const roundedRating = Math.round(rating * 2) / 2;
  const fullStars = Math.floor(roundedRating);
  const hasHalfStar = roundedRating % 1 !== 0;
  const emptyStars = 5 - fullStars - (hasHalfStar ? 1 : 0);

  for (let i = 0; i < fullStars; i++)
    stars.push(
      <img
        key={`full-${i}`}
        src="/icons/full-star.png"
        alt="★"
        className={styles.star}
      />
    );
  if (hasHalfStar)
    stars.push(
      <img
        key="half"
        src="/icons/half-star.png"
        alt="☆"
        className={styles.star}
      />
    );
  for (let i = 0; i < emptyStars; i++)
    stars.push(
      <img
        key={`empty-${i}`}
        src="/icons/empty-star.png"
        alt="✩"
        className={styles.star}
      />
    );

  return <div className={styles.rating}>{stars}</div>;
};
import.meta.env.VITE_YANDEX_API_KEY;
const ApartmentCarousel = ({ photos, title }) =>
  photos.length > 0 && (
    <Carousel
      showThumbs={false}
      infiniteLoop
      useKeyboardArrows
      className={styles.carousel}
      showStatus={false}
    >
      {photos.map((photo) => (
        <div
          key={photo.id}
          className={styles.slideWrapper}
          style={{
            "--bg-image": `url(${import.meta.env.VITE_BACKEND_URL}${
              photo.url
            })`,
          }}
        >
          <img
            src={`${import.meta.env.VITE_BACKEND_URL}${photo.url}`}
            alt={`Фото квартиры ${title}`}
            className={styles.carouselImg}
          />
        </div>
      ))}
    </Carousel>
  );

const ApartmentDetails = ({
  apartment,
  title,
  rent,
  deposit,
  rental_type,
  status,
}) => {
  const formatRentalType = (type) =>
    ({ monthly: "Помесячно", daily: "Посуточно" }[type] || type);

  return (
    <>
      <h2 className={styles.title}>
        {" "}
        {apartment.city}, {apartment.street} {apartment.building}
      </h2>
      <h2 className={styles.price}> {rent.toLocaleString()} ₽/мес</h2>
      <section className={styles.address}>
        <p>
          Этаж: {apartment.floor} | Комнат: {apartment.rooms}
        </p>
        <p>
          Тип дома: {apartment.construction_type} ({apartment.construction_year}{" "}
          г.)
        </p>
        <p>Ремонт: {apartment.remont}</p>
      </section>
      <section className={styles.details}>
        <p>Залог: {deposit.toLocaleString()} ₽</p>
        <p>Тип аренды: {formatRentalType(rental_type)}</p>
        <p>Статус: {status}</p>
      </section>
    </>
  );
};

const Amenities = (props) => {
  const featureIcons = {
    pets: { label: "Можно с животными", src: "/icons/pets.svg" },
    babies: { label: "Можно с детьми", src: "/icons/baby.svg" },
    smoking: { label: "Можно курить", src: "/icons/smoking.svg" },
    internet: { label: "Интернет", src: "/icons/internet.svg" },
    washing_machine: { label: "Стиралка", src: "/icons/washing_machine.svg" },
    tv: { label: "Телевизор", src: "/icons/tv.svg" },
    conditioner: { label: "Кондиционер", src: "/icons/conditioner.svg" },
    dishwasher: { label: "Посудомойка", src: "/icons/dishwasher.svg" },
    concierge: { label: "Консьерж", src: "/icons/concierge.svg" },
    garbage_chute: { label: "Мусоропровод", src: "/icons/garbage.svg" },
    elevator: { label: "Лифт", src: "/icons/elevator.svg" },
  };

  return (
    <section className={styles.features}>
      <h2 className={styles.title}>Удобства</h2>
      <div className={styles.featuresGrid}>
        {Object.entries(featureIcons).map(([key, { label, src }]) => (
          <FeatureCardStatic
            key={key}
            label={label}
            icon={<img src={src} alt={label} className={styles.featureIcon} />}
            selected={props[key]}
          />
        ))}
      </div>
    </section>
  );
};
const OwnerCard = ({ owner, onClick }) => {
  if (!owner) return <p>Загрузка владельца...</p>;
  return (
    <div className={styles.ownerCard}>
      <div
        className={styles.ownerInfo}
        onClick={onClick}
        style={{ cursor: "pointer" }}
      >
        <img
          src={owner.profile_picture || "/images/no-photo.png"}
          alt={`${owner.name} ${owner.surname}`}
          className={styles.avatar}
        />

        <div className={styles.profileInfo}>
          <div className={styles.ownerName}>
            <p>
              {owner.name} {owner.surname}
            </p>
          </div>
          <Stars rating={owner.rating} />
        </div>
      </div>
      <button className={styles.chatButton}>Написать в чат</button>
    </div>
  );
};

const AdvertPage = () => {
  const { id } = useParams();
  const [advert, setAdvert] = useState(null);
  const [owner, setOwner] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [mapMarker, setMapMarker] = useState([55.75396, 37.620393]);
  const [mapCenter, setMapCenter] = useState([55.75396, 37.620393]);
  const [similarAdverts, setSimilarAdverts] = useState([]);
  useEffect(() => {
    const loadAdvert = async () => {
      try {
        setLoading(true);
        const data = await fetchAdvertById(id);
        setAdvert(data);

        const userResponse = await api.get(`/users/${data.apartment.user_id}`);
        setOwner(userResponse.data);
        if (data.apartment.latitude && data.apartment.longitude) {
          setMapMarker([data.apartment.latitude, data.apartment.longitude]);
          setMapCenter([data.apartment.latitude, data.apartment.longitude]);
        }
        const similar = await fetchSimilarAdverts(data);
        setSimilarAdverts(similar);
      } catch (err) {
        console.error("Ошибка при загрузке:", err);
        setError("Ошибка при загрузке объявления");
      } finally {
        setLoading(false);
      }
    };
    loadAdvert();
  }, [id]);

  if (loading)
    return (
      <div>
        <SubNavbar />
        <Spinner />
      </div>
    );
  if (error) return <p className={styles.error}>{error}</p>;
  if (!advert) return null;

  const {
    apartment,
    title,
    rent,
    deposit,
    rental_type,
    pets,
    babies,
    smoking,
    internet,
    washing_machine,
    tv,
    conditioner,
    dishwasher,
    concierge,
    status,
  } = advert;
  const photos = apartment.apartment_photos || [];

  return (
    <>
      <SubNavbar />
      <div className={styles.container}>
        <div className={styles.top}>
          <div className={styles.left}>
            <ApartmentCarousel photos={photos} title={title} />
            <div className={styles.advertInfo}>
              {" "}
              <ApartmentDetails
                apartment={apartment}
                title={title}
                rent={rent}
                deposit={deposit}
                rental_type={rental_type}
                status={status}
              />
              <Amenities
                apartment={apartment}
                pets={pets}
                babies={babies}
                smoking={smoking}
                internet={internet}
                washing_machine={washing_machine}
                tv={tv}
                conditioner={conditioner}
                dishwasher={dishwasher}
                concierge={concierge}
              />
            </div>
          </div>
          <div className={styles.right}>
            <OwnerCard
              owner={owner}
              onClick={() => owner && setShowModal(true)}
            />

            {showModal && owner && (
              <OwnerModal
                ownerId={owner.id}
                onClose={() => setShowModal(false)}
              />
            )}
            <div className={styles.mapContainer}>
              <MapGLForm
                center={mapCenter}
                markerPosition={mapMarker}
                onSelect={({ lng, lat }) => {
                  setForm((prev) => ({
                    ...prev,
                    latitude: lat,
                    longitude: lng,
                  }));
                  setMapCenter([lat, lng]);
                  setMapMarker([lat, lng]);
                }}
              />
            </div>
          </div>
        </div>
        <div className={styles.bottom}>
          <h2 className={styles.similarTitle}>Похожие объявления</h2>
          <div className={styles.similarList}>
            {similarAdverts.length > 0 ? (
              similarAdverts.map((adv) => (
                <div key={adv.id} className={styles.similarCardWrapper}>
                  <AdvertCard advert={adv} />
                </div>
              ))
            ) : (
              <p className={styles.noSimilar}>Нет похожих объявлений</p>
            )}
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default AdvertPage;
