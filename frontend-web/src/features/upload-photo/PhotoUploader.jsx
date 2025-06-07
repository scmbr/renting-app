import { useState } from "react";
import styles from "./PhotoUploader.module.css";

export default function PhotoUploader({ previewUrls, setPreviewUrls }) {
  const [isDragging, setIsDragging] = useState(false);

  const handleFiles = (files) => {
    const fileArray = Array.from(files).filter((file) =>
      file.type.startsWith("image/")
    );
    if (fileArray.length === 0) return;

    const newUrls = fileArray.map((file) => URL.createObjectURL(file));
    setPreviewUrls((prev) => [...prev, ...newUrls]);
  };

  const handleDragOver = (e) => {
    e.preventDefault();
    setIsDragging(true);
  };

  const handleDragLeave = (e) => {
    e.preventDefault();
    setIsDragging(false);
  };

  const handleDrop = (e) => {
    e.preventDefault();
    setIsDragging(false);
    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      handleFiles(e.dataTransfer.files);
      e.dataTransfer.clearData();
    }
  };

  const handleInputChange = (e) => {
    handleFiles(e.target.files);
  };

  // Новая функция удаления по индексу
  const handleRemovePhoto = (index) => {
    setPreviewUrls((prev) => prev.filter((_, i) => i !== index));
  };

  return (
    <div className={styles.section}>
      <div
        className={`${styles.dropZone} ${isDragging ? styles.dragActive : ""}`}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        onDrop={handleDrop}
      >
        <label className={styles.fileInputLabel}>
          <span>Перетащите фото сюда или нажмите для выбора</span>
          <input
            type="file"
            multiple
            accept="image/*"
            onChange={handleInputChange}
            className={styles.fileInputHidden}
          />
        </label>
      </div>
      <div className={styles.previewContainer}>
        {previewUrls.map((url, idx) => (
          <div
            key={idx}
            className={styles.previewWrapper}
            onClick={() => handleRemovePhoto(idx)}
            title="Кликните чтобы удалить фото"
            style={{
              cursor: "pointer",
              position: "relative",
              display: "inline-block",
            }}
          >
            <img
              src={url}
              alt={`preview-${idx}`}
              className={styles.previewImage}
            />
            <span className={styles.deleteIcon}>×</span>
          </div>
        ))}
      </div>
    </div>
  );
}
