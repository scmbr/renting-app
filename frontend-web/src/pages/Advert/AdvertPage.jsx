import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { fetchAdvertById } from "@/entities/advert/model";
import styles from "./AdvertPage.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import api from "@/shared/api/axios";
import { Carousel } from "react-responsive-carousel";
import "react-responsive-carousel/lib/styles/carousel.min.css";
import OwnerModal from "@/widgets/OwnerModal/OwnerModal.jsx";
const AdvertPage = () => {
  const { id } = useParams();
  const [advert, setAdvert] = useState(null);
  const [photos, setPhotos] = useState([]);
  const [owner, setOwner] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [showModal, setShowModal] = useState(false);
  useEffect(() => {
    const loadAdvert = async () => {
      try {
        setLoading(true);
        const data = await fetchAdvertById(id);
        setAdvert(data);

        const photosResponse = await api.get(
          `/apartment/${data.apartment.id}/photos`
        );
        setPhotos(photosResponse.data || []);

        const userResponse = await api.get(`/users/${data.apartment.user_id}`);
        setOwner(userResponse.data);
      } catch (err) {
        console.error("Ошибка при загрузке:", err);
        setError("Ошибка при загрузке объявления");
      } finally {
        setLoading(false);
      }
    };

    loadAdvert();
  }, [id]);

  if (loading) return <p>Загрузка объявления...</p>;
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
  } = advert;
  const formatRentalType = (type) => {
    switch (type) {
      case "monthly":
        return "Помесячно";
      case "daily":
        return "Посуточно";
      default:
        return type;
    }
  };
  return (
    <>
      <SubNavbar />
      <div className={styles.pageWrapper}>
        <div className={styles.container}>
          {photos.length > 0 && (
            <div className={styles.carouselWrapper}>
              <Carousel
                className={styles.carousel}
                showThumbs={false}
                dynamicHeight={false}
                infiniteLoop
                useKeyboardArrows
                autoFocus={false}
              >
                {photos.map((photo) => (
                  <div key={photo.id}>
                    <img src={photo.url} alt={`Фото квартиры ${title}`} />
                  </div>
                ))}
              </Carousel>
            </div>
          )}
          <div className={styles.info}>
            <h2 className={styles.title}>{advert.title}</h2>

            <section className={styles.address}>
              <p>
                Адрес: {apartment.city}, {apartment.street} {apartment.building}
              </p>
              <p>
                Этаж: {apartment.floor} | Комнат: {apartment.rooms}
              </p>
              <p>
                Тип дома: {apartment.construction_type} (
                {apartment.construction_year} г.)
              </p>
              <p>Ремонт: {apartment.remont}</p>
            </section>

            <section className={styles.details}>
              <p>
                Аренда:{" "}
                <span className={styles.price}>
                  {rent.toLocaleString()} ₽/мес
                </span>
              </p>
              <p>Залог: {deposit.toLocaleString()} ₽</p>
              <p>Тип аренды: {formatRentalType(rental_type)}</p>
              <p>Статус: {advert.status}</p>
            </section>

            <section className={styles.features}>
              <h2 className={styles.title}>Удобства</h2>
              <ul>
                <li>Питомцы: {pets ? "разрешены" : "запрещены"}</li>
                <li>Дети: {babies ? "разрешены" : "запрещены"}</li>
                <li>Курение: {smoking ? "разрешено" : "запрещено"}</li>
                <li>Интернет: {internet ? "есть" : "нет"}</li>
                <li>Стиральная машина: {washing_machine ? "есть" : "нет"}</li>
                <li>Телевизор: {tv ? "есть" : "нет"}</li>
                <li>Кондиционер: {conditioner ? "есть" : "нет"}</li>
                <li>Посудомоечная машина: {dishwasher ? "есть" : "нет"}</li>
                <li>Консьерж: {concierge ? "есть" : "нет"}</li>
                <li>
                  Мусоропровод: {apartment.garbage_chute ? "есть" : "нет"}
                </li>
                <li>Лифт: {apartment.elevator ? "есть" : "нет"}</li>
                <li>Тип ванной: {apartment.bathroom_type}</li>
              </ul>
            </section>
          </div>
        </div>

        <div
          className={styles.ownerCard}
          onClick={() => setShowModal(true)}
          style={{ cursor: "pointer" }}
        >
          {owner ? (
            <>
              {owner.profile_picture && (
                <img
                  src={owner.profile_picture}
                  alt={`${owner.name} ${owner.surname}`}
                  className={styles.avatar}
                />
              )}
              <div className={styles.ownerName}>
                <p>
                  {owner.name} {owner.surname}
                </p>
              </div>
              <p>
                <strong>Email:</strong> {owner.email}
              </p>
              <p>
                <strong>Телефон:</strong> {owner.phone}
              </p>
              <p>
                <strong>Рейтинг:</strong>
              </p>
              <div className={styles.rating}>{renderStars(owner.rating)}</div>
            </>
          ) : (
            <p>Загрузка владельца...</p>
          )}

          {showModal && (
            <OwnerModal
              ownerId={owner.id}
              onClose={() => setShowModal(false)}
            />
          )}
        </div>
      </div>
    </>
  );
};

export default AdvertPage;
const renderStars = (rating) => {
  const stars = [];
  const fullStars = Math.floor(rating);
  const hasHalfStar = rating % 1 >= 0.25 && rating % 1 <= 0.75;
  const emptyStars = 5 - fullStars - (hasHalfStar ? 1 : 0);

  for (let i = 0; i < fullStars; i++) {
    stars.push(
      <img
        key={`full-${i}`}
        src="/icons/full-star.png"
        alt="★"
        className={styles.star}
      />
    );
  }

  if (hasHalfStar) {
    stars.push(
      <img
        key="half"
        src="/icons/half-star.png"
        alt="☆"
        className={styles.star}
      />
    );
  }

  for (let i = 0; i < emptyStars; i++) {
    stars.push(
      <img
        key={`empty-${i}`}
        src="/icons/empty-star.png"
        alt="✩"
        className={styles.star}
      />
    );
  }

  return stars;
};
