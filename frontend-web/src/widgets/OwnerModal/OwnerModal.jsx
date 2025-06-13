import React, { useEffect, useState } from "react";
import api from "@/shared/api/axios";
import styles from "./OwnerModal.module.css";

const OwnerModal = ({ ownerId, onClose }) => {
  const [owner, setOwner] = useState(null);
  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState("");
  const [rating, setRating] = useState(0);
  const [hoverRating, setHoverRating] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchOwnerData = async () => {
      try {
        setLoading(true);
        const ownerResponse = await api.get(`/users/${ownerId}`);
        setOwner(ownerResponse.data);

        const commentsResponse = await api.get(`/users/${ownerId}/reviews`);
        setComments(commentsResponse.data || []);
      } catch (err) {
        setError("Ошибка загрузки данных пользователя");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchOwnerData();
  }, [ownerId]);

  const handleAddComment = async () => {
    if (!newComment.trim() || rating === 0) return;

    try {
      const payload = {
        target_id: ownerId,
        rating,
        comment: newComment.trim(),
      };

      const res = await api.post("/reviews", payload);
      setComments((prev) => [...prev, res.data]);
      setNewComment("");
      setRating(0);
      setHoverRating(0);
    } catch (err) {
      console.error("Ошибка добавления комментария:", err);
      alert("Не удалось отправить комментарий. Попробуйте позже.");
    }
  };

  const renderStars = () => {
    return [...Array(5)].map((_, i) => {
      const starIndex = i + 1;
      const isFilled = hoverRating
        ? starIndex <= hoverRating
        : starIndex <= rating;
      return (
        <img
          key={starIndex}
          src={isFilled ? "/icons/full-star.png" : "/icons/empty-star.png"}
          alt={isFilled ? "Полная звезда" : "Пустая звезда"}
          className={styles.star}
          onClick={() => setRating(starIndex)}
          onMouseEnter={() => setHoverRating(starIndex)}
          onMouseLeave={() => setHoverRating(0)}
          style={{ cursor: "pointer" }}
        />
      );
    });
  };

  if (loading) return <p>Загрузка данных...</p>;
  if (error) return <p className={styles.error}>{error}</p>;
  if (!owner) return null;

  return (
    <div className={styles.backdrop} onClick={onClose}>
      <div className={styles.modal} onClick={(e) => e.stopPropagation()}>
        <button className={styles.closeButton} onClick={onClose}>
          ×
        </button>
        <h2>
          {owner.name} {owner.surname}
        </h2>
        <p>Email: {owner.email}</p>
        <p>Телефон: {owner.phone}</p>

        <div className={styles.section}>
          <h3>Комментарии</h3>
          {comments.length > 0 ? (
            <ul className={styles.commentList}>
              {comments.map((comment) => (
                <li key={comment.id} className={styles.commentItem}>
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
                        {renderStarsStatic(comment.rating)}
                      </div>
                    </div>
                  </div>
                  <p className={styles.commentText}>{comment.comment}</p>
                </li>
              ))}
            </ul>
          ) : (
            <p>Пока нет комментариев</p>
          )}
        </div>

        <div className={styles.addCommentSection}>
          <h3>Оставить комментарий</h3>
          <div className={styles.starSelector}>{renderStars()}</div>
          <textarea
            value={newComment}
            onChange={(e) => setNewComment(e.target.value)}
            placeholder="Оставьте комментарий..."
            className={styles.textarea}
          />
          <button onClick={handleAddComment} className={styles.submitButton}>
            Отправить
          </button>
        </div>
      </div>
    </div>
  );
};

const renderStarsStatic = (rating) => {
  return [...Array(5)].map((_, i) => {
    const starIndex = i + 1;
    const isFilled = starIndex <= rating;
    return (
      <img
        key={starIndex}
        src={isFilled ? "/icons/full-star.png" : "/icons/empty-star.png"}
        alt={isFilled ? "Полная звезда" : "Пустая звезда"}
        className={styles.starStatic}
      />
    );
  });
};

export default OwnerModal;
