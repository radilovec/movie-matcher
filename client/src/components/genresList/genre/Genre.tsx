import style from "./style.module.css";

type Props = {
  name: string;
};

export const Genre: React.FC<Props> = ({ name }) => {
  return <span className={style.genre}>{name}</span>;
};
