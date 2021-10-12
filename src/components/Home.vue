<template>
  <div id="home">
    <h2 class="sub-title">フリー素材サイト</h2>
    <p class="home-text">Home page</p>
    <br /><br />

    <h3 class="members-title">プロジェクトメンバー</h3>
    <p class="members">
      メインプログラマー : TA吉村<br />
      ちょこっとプログラマー : TA赤木<br />
      企画発案 : TA綾井<br />
      企画 : TAサイタ, TA本多<br />
    </p>

    <div id="b">
      <button
        class="skew-button"
        onclick="location.href='https://linkk.la/genkiakagi'"
      >
        <span>bonus!!</span>
      </button>
    </div>
  </div>
</template>


<script>
export default {};
</script>

<style>
#home {
  width: 70%;
  margin: auto;
}
.home-text {
  text-align: center;
}
.sub-title {
  margin-top: 10%;
  text-align: center;
}
.members-title {
  margin-top: 60%;
  text-align: left;
}
.members {
  text-align: left;
}

.skew-button {
  margin-left: 35%;
}

* {
  box-sizing: border-box;
}

body {
  height: 100vh;
  background-color: #fff;
}

.skew-button {
  position: absolute;
  left: 50%;
  top: 50%;
}

.skew-button {
  cursor: pointer;
  outline: none;
  transform: translate3d(-50%, -50%, 0);
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
  position: relative;
  display: inline-block;
  padding: 0;
  background-color: transparent;
  border: none;
  font: 1.3rem "Anton", sans-serif;
  text-transform: uppercase;
  letter-spacing: 0.1em;
  color: #fff;
}
.skew-button span {
  transform: translate3d(0, 0, 0);
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
  position: relative;
  display: block;
  width: 100%;
  height: 100%;
  padding: 20px 50px;
  z-index: 3;
}
.skew-button:before,
.skew-button span:before {
  content: "";
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%) skewX(-20deg);
  width: 100%;
  height: 50px;
  border: 1px solid #0e3876;
  z-index: 2;
}
.skew-button span:before {
  display: block;
  transform: translate(-50%, -50%) skew(-20deg);
  transform-origin: center center;
  z-index: -1;
}
.skew-button:after {
  transition: all 100ms ease-out;
  content: "";
  position: absolute;
  left: 0;
  top: 0;
  transform: skewX(-20deg) scaleX(0);
  width: 100%;
  height: 100%;
  background-color: #0e3876;
  border: 1px solid transparent;
  z-index: 1;
}
.skew-button:hover span:before {
  -webkit-animation: fill 0.5s ease-out;
  animation: fill 0.5s ease-out;
  -webkit-animation-fill-mode: forwards;
  animation-fill-mode: forwards;
}
.skew-button:hover:after {
  -webkit-animation: swoosh 0.7s ease-in;
  animation: swoosh 0.7s ease-in;
}
.skew-button:active:before {
  background-color: #030d1b;
}

@-webkit-keyframes swoosh {
  0% {
    transform-origin: left center;
  }
  20% {
    transform: skewX(-20deg) scaleX(1);
    transform-origin: left center;
  }
  21% {
    transform-origin: right center;
  }
  50% {
    transform: skewX(-20deg) scaleX(0);
    transform-origin: right center;
  }
}

@keyframes swoosh {
  0% {
    transform-origin: left center;
  }
  20% {
    transform: skewX(-20deg) scaleX(1);
    transform-origin: left center;
  }
  21% {
    transform-origin: right center;
  }
  50% {
    transform: skewX(-20deg) scaleX(0);
    transform-origin: right center;
  }
}
@-webkit-keyframes fill {
  0% {
    background-color: rgba(14, 56, 118, 0);
  }
  20% {
    opacity: 0.5;
    border-width: 15px;
    border-color: #092248;
  }
  100% {
    opacity: 0.5;
    border-width: 1px;
    background-color: rgba(14, 56, 118, 0.5);
  }
}
@keyframes fill {
  0% {
    background-color: rgba(14, 56, 118, 0);
  }
  20% {
    opacity: 0.5;
    border-width: 15px;
    border-color: #092248;
  }
  100% {
    opacity: 0.5;
    border-width: 1px;
    background-color: rgba(14, 56, 118, 0.5);
  }
}
</style>
