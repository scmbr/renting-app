import { useEffect, useState } from "react";
import { fetchFavorites } from "@/entities/favorites/api";
import { fetchAdvertById } from "@/entities/advert/model";
import FavoritesList from "@/widgets/FavoritesList/FavoritesList.jsx";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";
import styles from "./FavoritesPage.module.css";
const FavoritesPage = () => {
  const [favoriteAdverts, setFavoriteAdverts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    async function loadFavorites() {
      setLoading(true);
      setError(null);
      try {
        const favorites = await fetchFavorites();
        const advertsPromises = favorites.map((fav) =>
          fetchAdvertById(fav.advert_id)
        );
        const adverts = await Promise.all(advertsPromises);
        setFavoriteAdverts(adverts);
      } catch (err) {
        setError("Ошибка при загрузке избранного");
        console.error(err);
      } finally {
        setLoading(false);
      }
    }

    loadFavorites();
  }, []);

  const handleRemoveFromFavorites = (removedId) => {
    setFavoriteAdverts((prev) =>
      prev.filter((advert) => advert.id !== removedId)
    );
  };

  if (loading) return <p>Загрузка избранного...</p>;
  if (favoriteAdverts.length === 0)
    return (
      <>
        <SubNavbar />
        <NavPanel />
        <div className={styles.container}>
          <h1 className={styles.title}>Избранное</h1>
        </div>
      </>
    );

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.container}>
        <h1 className={styles.title}>Избранное</h1>
        <FavoritesList
          adverts={favoriteAdverts}
          total={favoriteAdverts.length}
          onRemoveFavorite={handleRemoveFromFavorites}
        />
      </div>
    </>
  );
};

export default FavoritesPage;
