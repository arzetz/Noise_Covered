import React from "react";
import "./AboutUs.css";
import { useState, useEffect } from "react";
import { Container } from "react-bootstrap";
function AboutUs() {
  const [isFirstEffectDone, setIsFirstEffectDone] = useState(false);
  const [isSecondEffectDone, setIsSecondEffectDone] = useState(false);
  const [isThirdEffectDone, setIsThirdEffectDone] = useState(false);
  const [displayedFirstText, setDisplayedFirstText] = useState("");
  const [displayedSecondText, setDisplayedSecondText] = useState("");
  const [displayedThirdText, setDisplayedThirdText] = useState("");
  const [displayedFourthText, setDisplayedFourthText] = useState("");
  const [isFading, setIsFading] = useState(false);
  const [isHidden, setIsHidden] = useState(false);
  const [time, setTime] = useState(50);
  const firstText = "К то мы?";
  const secondText = `В  2024 году был создан розничный магазин пластинок под названием NOISE COVERED.
Вдохновлённая идеей профессиональной записи каверов, наша команда
нашла большое количество артистов, кто мог бы заниматься созданием исполнений
и их ресейлом вместе с нами.

Данная площадка - место собрания уникальной коллекции творческих
людей, стать частью которой вы можете просто  приобретя интересующую вас
запись или выставив на продажу свой товар.
`;
const thirdText = `\nХотите развиваться с нами?`;

useEffect(() => {
  let currentIndex = 0;
  let interval;

  const delayTimeout = setTimeout(() => {
    interval = setInterval(() => {
      if (currentIndex < firstText.length-1) {
        setDisplayedFirstText((prev) => prev + firstText[currentIndex]);
        currentIndex++;
      } else {
        clearInterval(interval);
        setTimeout(() => {
          setIsFading(true);
          setTimeout(() => {
          setIsFirstEffectDone(true);
          setIsHidden(true)},2000)
        }, 2000);
      }
    }, time);
  }, 1000);

  return () => {
    clearInterval(interval);
    clearTimeout(delayTimeout);
  };
}, [time]);

useEffect(() => {
  if (isFirstEffectDone) {
    setTime(25)
    let currentIndex = 0;
    let interval;

      interval = setInterval(() => {
        if (currentIndex < secondText.length-1) {
          setDisplayedSecondText((prev) => prev + secondText[currentIndex]);
          currentIndex++;
        } else {
          clearInterval(interval);
          setIsSecondEffectDone(true)
        }
      }, time);

    return () => {
      clearInterval(interval);
    };
  }
}, [isFirstEffectDone, time]);

useEffect(() => {
  if (isSecondEffectDone) {

    let currentIndex = 0;
    let interval;

      interval = setInterval(() => {
        if (currentIndex < thirdText.length-1) {
          setDisplayedThirdText((prev) => prev + thirdText[currentIndex]);
          currentIndex++;
        } else {
          clearInterval(interval);
          setIsThirdEffectDone(true)
        }
      }, time);

    return () => {
      clearInterval(interval);
    };
  }
}, [isSecondEffectDone, time]);

  return (
    <div className="background-container_about_us">
      <Container className={`${isHidden ? "about_us_container" : "about_us_container_centered"}`}>
      <span className={`custom-font1 ${isFading ? "fade-out" : ""} ${isHidden ? "fade-out-hidden" : ""}`}>{displayedFirstText}</span>
      <span className="custom-font2">{displayedSecondText}</span>
      <br></br>
      <span className="custom-font3">{displayedThirdText}</span>
        </Container>
    </div>
  );
}
export default AboutUs;
