import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "@/shared/api/axios";
import styles from "./MyProfile.module.css";
import SubNavbar from "@/widgets/SubNavbar/SubNavbar.jsx";
import NavPanel from "@/widgets/NavPanel/NavPanel.jsx";

const MyProfile = () => {
  const [profile, setProfile] = useState(null);
  const [comments, setComments] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const profileRes = await api.get("/me");
        const profileData = profileRes.data;
        setProfile(profileData);

        const commentsRes = await api.get(`/users/${profileData.id}/reviews`);
        setComments(commentsRes.data || []);
      } catch (err) {
        if (err?.response?.status === 401) {
          navigate("/login");
        } else {
          setError("Не удалось загрузить данные профиля");
        }
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [navigate]);

  if (loading) return <p className={styles.loading}>Загрузка профиля...</p>;
  if (error) return <p className={styles.error}>{error}</p>;

  return (
    <>
      <SubNavbar />
      <NavPanel />
      <div className={styles.container}>
        <h1 className={styles.title}>Мой профиль</h1>
        <div className={styles.profileCard}>
          {profile.profile_picture && (
            <img
              src={profile.profile_picture}
              alt="Аватар"
              className={styles.avatar}
            />
          )}
          <span>
            {profile.name} {profile.surname}
          </span>
          <div className={styles.rating}>{renderStars(profile.rating)}</div>
        </div>

        <div className={styles.commentList}>
          <h2 className={styles.commentTitle}>Отзывы</h2>
          {comments.length === 0 ? (
            <p className={styles.noComments}>Пока нет отзывов</p>
          ) : (
            comments.map((comment, idx) => (
              <div key={idx} className={styles.comment}>
                <div className={styles.commentHeader}>
                  <div className={styles.commentAuthorInfo}>
                    <img
                      src={
                        comment.author?.profile_picture || "/default-avatar.png"
                      }
                      alt="автор"
                      className={styles.commentAuthorAvatar}
                    />
                    <div className={styles.commentAuthorText}>
                      <div className={styles.commentAuthorName}>
                        {comment.author?.name} {comment.author?.surname}
                      </div>
                      <div className={styles.commentRating}>
                        {renderStars(comment.rating)}
                      </div>
                    </div>
                  </div>
                </div>
                <p>{comment.comment}</p>
              </div>
            ))
          )}
        </div>
      </div>
    </>
  );
};

export default MyProfile;

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
