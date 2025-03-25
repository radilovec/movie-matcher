import styles from "./style.module.css";
import { MdErrorOutline } from "react-icons/md";

interface Props {
  image: string;
  alt: string;
  type?: "circle" | "rectangle";
}

export const Image: React.FC<Props> = ({ type = "rectangle", image, alt }) => {
  return (
    <div className={`${type === "rectangle" ? styles.wrapper : styles.wrapperCircle}`}>
      {image ? (
        <img
          src={image}
          className={`${type === "rectangle" ? styles.image : styles.imgCircle}`}
          alt={alt}
        />
      ) : (
        <div className={styles.box}>
          <MdErrorOutline size={50} />
        </div>
      )}
    </div>
  );
};
