import styles from "./ConfirmModal.module.css";

const ConfirmModal = ({ isOpen, onConfirm, onCancel, message }) => {
  if (!isOpen) return null;

  return (
    <div className={styles.overlay}>
      <div className={styles.modal}>
        <p>{message}</p>
        <div className={styles.buttons}>
          <button className={styles.confirm} onClick={onConfirm}>
            Удалить
          </button>
          <button className={styles.cancel} onClick={onCancel}>
            Отмена
          </button>
        </div>
      </div>
    </div>
  );
};

export default ConfirmModal;
