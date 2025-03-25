import { Genre } from "./genre/Genre";
import style from "./style.module.css";

type GenreType = {
  id: number;
  name: string;
};

type Props = {
  genres: GenreType[];
};

export const GenresList: React.FC<Props> = ({ genres }) => {
  return (
    <div className={style.list}>
      {genres.map((genre) => (
        <Genre key={genre.id} name={genre.name} />
      ))}
    </div>
  );
};
