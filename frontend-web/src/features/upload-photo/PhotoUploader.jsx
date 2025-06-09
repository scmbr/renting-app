import { useState } from "react";
import styles from "./PhotoUploader.module.css";

export default function PhotoUploader({
  previewUrls,
  setPreviewUrls,
  photos,
  setPhotos,
  setPhotosToDelete,
}) {
  const [isDragging, setIsDragging] = useState(false);

  const handleFiles = (files) => {
    const fileArray = Array.from(files).filter((file) =>
      file.type.startsWith("image/")
    );
    if (fileArray.length === 0) return;

    const newUrls = fileArray.map((file) => ({
      id: null,
      url: URL.createObjectURL(file),
    }));
    setPreviewUrls((prev) => [...prev, ...newUrls]);
    setPhotos((prev) => [...prev, ...fileArray]);
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

  const handleRemovePhoto = (index) => {
    const photoToRemove = previewUrls[index];
    if (photoToRemove.id) {
      setPhotosToDelete((prev) => [...prev, photoToRemove.id]);
    } else {
      setPhotos((prev) => {
        return prev.filter((_, i) => {
          const oldPhotosCount = previewUrls.filter((p) => p.id).length;
          return i !== index - oldPhotosCount;
        });
      });
    }

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
        {previewUrls.map((photo, idx) => (
          <div
            key={photo.id ?? photo.url}
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
              src={photo.url}
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
