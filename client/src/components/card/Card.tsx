import React, { JSX, useState } from "react";
import style from "./style.module.css";
import { Image } from "../image/Image";
import { GenresList } from "../genresList/GenresList";
import { FaStar, FaStarHalf, FaInfoCircle } from "react-icons/fa";
import { InfBox } from "./infBox/InfBox";

const genres = [
  {
    id: 1,
    name: "thriller",
  },
  {
    id: 2,
    name: "drama",
  },
  {
    id: 3,
    name: "horror",
  },
  {
    id: 4,
    name: "thriller",
  },
  {
    id: 5,
    name: "criminal roman",
  },
  {
    id: 6,
    name: "horror",
  },
];

const rating = 3.12;

const countStars = (rating: number): JSX.Element[] => {
  const roundedRating = Math.round(rating * 2) / 2;

  const fullStars = Math.floor(roundedRating);
  const hasHalfStar = roundedRating % 1 !== 0;
  const emptyStars = 5 - fullStars - (hasHalfStar ? 1 : 0);

  const stars: JSX.Element[] = [];

  for (let i = 0; i < fullStars; i++) {
    stars.push(<FaStar key={`full-${i}`} color="gold" />);
  }

  if (hasHalfStar) {
    stars.push(<FaStarHalf key="half" color="gold" />);
  }

  for (let i = 0; i < emptyStars; i++) {
    stars.push(<FaStar key={`empty-${i}`} color="gray" />);
  }

  return stars;
};

export const Card: React.FC = () => {
  const [isInfoVisible, setIsInfoVisible] = useState(false);

  const toggleInfo = () => {
    setIsInfoVisible(!isInfoVisible);
  };

  return (
    <div className={style.card}>
      <FaInfoCircle className={style.showBtn} onClick={toggleInfo} style={{ cursor: "pointer" }} />
      <Image image={"/img/poster.jpg"} alt={"poster"} />
      <div className={style.inf}>
        <div className={style.infVisible}>
          <div className={style.genresBox}>
            <div className={style.image}>
              <Image image={"img/bg.jpg"} alt={"from"} type="circle" />
            </div>
            <GenresList genres={genres} />
          </div>
          <p className={style.name}>Seven</p>
          <div className={style.rating}>
            <div className={style.stars}>{countStars(rating)}</div>
            <span className={style.voteCount}>(19291)</span>
          </div>
        </div>

        <div className={`${style.infInvisible} ${isInfoVisible ? style.visible : ""}`}>
          <p className={style.desc}>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptates quas, dignissimos
            magnam officiis impedit ullam iusto, voluptatum, perferendis enim sit suscipit dolorem
            deserunt quae dicta vero dolor hic aperiam ducimus repellat quasi perspiciatis est
            explicabo porro. Quasi voluptatum nam recusandae facere modi? Quas alias aliquid
            facilis! Quos autem quas dolores!
          </p>

          <div className={style.moreInf}>
            <InfBox param="duration" val="149min" />
            <InfBox param="language" val="EN" />
            <InfBox param="released" val="1996" />
          </div>
        </div>
      </div>
    </div>
  );
};
