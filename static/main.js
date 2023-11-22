 // Theme switcher.
 function calculateSettingAsThemeString({ localStorageTheme, systemSettingDark }) {
  if (localStorageTheme !== null) {
    return localStorageTheme;
  }

  if (systemSettingDark.matches) {
    return "dark";
  }

  return "light";
}

const localStorageTheme = localStorage.getItem("theme");
const systemSettingDark = window.matchMedia("(prefers-color-scheme: dark)");

let currentThemeSetting = calculateSettingAsThemeString({ localStorageTheme, systemSettingDark });

console.log(`Current theme is ${currentThemeSetting}`);

// Target the button using the data attribute we added earlier.
const button = document.querySelector("[data-theme-toggle]");
const lightSwitchText = document.getElementById("light-switch-text");

document.querySelector("html").setAttribute("data-theme", currentThemeSetting);
const currentThemeIcon = currentThemeSetting === "dark" ? "🌞" : "🌚";
lightSwitchText.innerText = currentThemeIcon;
button.checked = currentThemeSetting === 'light';

button.addEventListener("click", () => {
  const newTheme = currentThemeSetting === "dark" ? "light" : "dark";

  // Update the button text.
  const newCta = newTheme === "dark" ? "🌞" : "🌚";
  lightSwitchText.innerText = newCta;  

  // Use an aria-label if you are omitting text on the button
  // and using sun/moon icons, for example.
  button.setAttribute("aria-label", newCta);

  // Update theme attribute on HTML to switch theme in CSS.
  document.querySelector("html").setAttribute("data-theme", newTheme);

  // Update in local storage.
  localStorage.setItem("theme", newTheme);

  // Update the currentThemeSetting in memory.
  currentThemeSetting = newTheme;
});