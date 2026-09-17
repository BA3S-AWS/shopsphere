const productsContainer = document.getElementById("products");
const cartDrawer = document.getElementById("cart-drawer");
const cartOverlay = document.getElementById("cart-overlay");
const openCartButton = document.getElementById("open-cart-button");
const closeCartButton = document.getElementById("close-cart-button");
const cartItemsContainer = document.getElementById("cart-items");
const cartCount = document.getElementById("cart-count");
const cartTotal = document.getElementById("cart-total");
const checkoutButton = document.getElementById("checkout-button");
const toast = document.getElementById("toast");
const searchInput = document.getElementById("search-input");

let userId = localStorage.getItem("shopsphere-user-id");

if (!userId) {
    userId = `user-${crypto.randomUUID()}`;
    localStorage.setItem("shopsphere-user-id", userId);
}

let products = [];
let productsById = {};


// ======================================================
// PANIER : OUVERTURE / FERMETURE
// ======================================================

function openCart() {
    cartDrawer.classList.add("open");
    cartOverlay.classList.add("open");
}

function closeCart() {
    cartDrawer.classList.remove("open");
    cartOverlay.classList.remove("open");
}

openCartButton.addEventListener("click", openCart);
closeCartButton.addEventListener("click", closeCart);
cartOverlay.addEventListener("click", closeCart);


// ======================================================
// TOAST
// ======================================================

function showToast(message, isError = false) {
    toast.textContent = message;

    toast.classList.remove("error");

    if (isError) {
        toast.classList.add("error");
    }

    toast.classList.add("show");

    setTimeout(() => {
        toast.classList.remove("show");
    }, 2500);
}


// ======================================================
// FORMAT PRIX
// ======================================================

function formatPrice(value) {
    return Number(value).toLocaleString("fr-FR", {
        minimumFractionDigits: 2,
        maximumFractionDigits: 2
    }) + " €";
}


// ======================================================
// NOM COMMERCIAL
// ======================================================

function displayName(product) {
    const name = product.name.toLowerCase();

    if (name.includes("laptop")) {
        return 'Laptop DevOps Pro 15"';
    }

    if (name.includes("clavier")) {
        return "Clavier Mécanique RGB";
    }

    if (name.includes("ecran") || name.includes("écran")) {
        return 'Écran 27" QHD';
    }

    return product.name;
}


// ======================================================
// ILLUSTRATIONS PRODUITS
// ======================================================

function getProductIllustration(product) {
    const name = product.name.toLowerCase();

    // LAPTOP
    if (name.includes("laptop")) {
        return `
            <svg viewBox="0 0 320 220" aria-label="Ordinateur portable">

                <defs>
                    <linearGradient
                        id="screenLaptop"
                        x1="0"
                        y1="0"
                        x2="1"
                        y2="1"
                    >
                        <stop offset="0%" stop-color="#162868"/>
                        <stop offset="50%" stop-color="#1f72e8"/>
                        <stop offset="100%" stop-color="#a24ddb"/>
                    </linearGradient>
                </defs>

                <rect
                    x="55"
                    y="28"
                    width="210"
                    height="135"
                    rx="10"
                    fill="#202532"
                />

                <rect
                    x="66"
                    y="39"
                    width="188"
                    height="112"
                    rx="4"
                    fill="url(#screenLaptop)"
                />

                <path
                    d="M70 138 L120 95 L150 120 L188 72 L248 145"
                    fill="none"
                    stroke="rgba(255,255,255,.65)"
                    stroke-width="5"
                />

                <path
                    d="M35 170 H285 L260 194 H60 Z"
                    fill="#9ca8bb"
                />

                <path
                    d="M105 177 H215"
                    stroke="#687386"
                    stroke-width="5"
                    stroke-linecap="round"
                />

            </svg>
        `;
    }

    // CLAVIER
    if (name.includes("clavier")) {
        return `
            <svg viewBox="0 0 320 220" aria-label="Clavier mécanique">

                <defs>
                    <linearGradient
                        id="keyboardBase"
                        x1="0"
                        y1="0"
                        x2="1"
                        y2="1"
                    >
                        <stop offset="0%" stop-color="#101722"/>
                        <stop offset="100%" stop-color="#30394a"/>
                    </linearGradient>
                </defs>

                <g transform="translate(28 62) skewX(-8)">

                    <rect
                        width="265"
                        height="108"
                        rx="13"
                        fill="url(#keyboardBase)"
                    />

                    ${createKeyboardKeys()}

                </g>

            </svg>
        `;
    }

    // ÉCRAN
    return `
        <svg viewBox="0 0 320 220" aria-label="Écran 27 pouces">

            <defs>
                <linearGradient
                    id="monitorScreen"
                    x1="0"
                    y1="0"
                    x2="1"
                    y2="1"
                >
                    <stop offset="0%" stop-color="#073f9b"/>
                    <stop offset="45%" stop-color="#287eee"/>
                    <stop offset="100%" stop-color="#f1729a"/>
                </linearGradient>
            </defs>

            <rect
                x="45"
                y="26"
                width="230"
                height="143"
                rx="9"
                fill="#232936"
            />

            <rect
                x="55"
                y="36"
                width="210"
                height="122"
                rx="3"
                fill="url(#monitorScreen)"
            />

            <path
                d="M57 143 L110 95 L145 120 L192 72 L262 148"
                fill="none"
                stroke="rgba(255,255,255,.65)"
                stroke-width="5"
            />

            <rect
                x="147"
                y="169"
                width="26"
                height="25"
                fill="#707989"
            />

            <rect
                x="110"
                y="192"
                width="100"
                height="9"
                rx="4"
                fill="#444c5a"
            />

        </svg>
    `;
}


function createKeyboardKeys() {
    let html = "";

    const colors = [
        "#38bdf8",
        "#8b5cf6",
        "#ec4899",
        "#22c55e",
        "#f97316"
    ];

    for (let row = 0; row < 4; row++) {
        for (let col = 0; col < 10; col++) {

            const x = 16 + col * 23;
            const y = 14 + row * 21;

            const color =
                colors[(row + col) % colors.length];

            html += `
                <rect
                    x="${x}"
                    y="${y}"
                    width="17"
                    height="14"
                    rx="2"
                    fill="#111827"
                    stroke="${color}"
                    stroke-width="1.6"
                />
            `;
        }
    }

    return html;
}


// ======================================================
// AFFICHAGE PRODUITS
// ======================================================

function renderProducts(list) {
    productsContainer.innerHTML = "";

    if (list.length === 0) {
        productsContainer.innerHTML =
            "<p>Aucun produit trouvé.</p>";

        return;
    }

    list.forEach((product, index) => {

        const card = document.createElement("article");

        card.className = "product-card";

        card.innerHTML = `
            <div class="product-image">

                ${
                    index === 0
                        ? '<span class="new-badge">Nouveau</span>'
                        : ""
                }

                ${getProductIllustration(product)}

            </div>

            <div class="product-content">

                <h3>
                    ${displayName(product)}
                </h3>

                <p class="product-description">
                    ${product.description || ""}
                </p>

                <div class="product-price">
                    ${formatPrice(product.price)}
                </div>

                <div class="stock">
                    📦 Stock : ${product.stock}
                </div>

                <button
                    class="add-cart-button"
                    type="button"
                    onclick="addToCart(${product.id})"
                >
                    🛒 Ajouter au panier
                </button>

            </div>
        `;

        productsContainer.appendChild(card);
    });
}


// ======================================================
// CHARGEMENT PRODUITS
// ======================================================

async function loadProducts() {
    try {

        const response =
            await fetch("/api/products");

        if (!response.ok) {
            throw new Error("Erreur produits");
        }

        products = await response.json();

        productsById = {};

        products.forEach(product => {
            productsById[Number(product.id)] = product;
        });

        renderProducts(products);

    } catch (error) {

        console.error(error);

        productsContainer.innerHTML =
            "<p>Impossible de charger le catalogue.</p>";
    }
}


// ======================================================
// RECHERCHE
// ======================================================

searchInput.addEventListener("input", () => {

    const query =
        searchInput.value.toLowerCase().trim();

    const filtered =
        products.filter(product => {

            const text = `
                ${product.name}
                ${product.description || ""}
                ${displayName(product)}
            `.toLowerCase();

            return text.includes(query);
        });

    renderProducts(filtered);
});


// ======================================================
// AJOUTER UN ARTICLE
// ======================================================

async function addToCart(productId) {

    const payload = {
        items: [
            {
                product_id: productId,
                quantity: 1
            }
        ]
    };

    try {

        const response =
            await fetch(`/api/cart/${userId}`, {

                method: "POST",

                headers: {
                    "Content-Type": "application/json"
                },

                body: JSON.stringify(payload)
            });

        if (!response.ok) {
            throw new Error("Erreur ajout panier");
        }

        showToast("Produit ajouté au panier ✓");

        await loadCart();

        openCart();

    } catch (error) {

        console.error(error);

        showToast(
            "Impossible d'ajouter le produit",
            true
        );
    }
}


// ======================================================
// CHARGEMENT DU PANIER
// ======================================================

async function loadCart() {
    try {

        const response =
            await fetch(`/api/cart/${userId}`);

        if (!response.ok) {
            renderEmptyCart();
            return;
        }

        const cart = await response.json();

        if (!cart.items || cart.items.length === 0) {
            renderEmptyCart();
            return;
        }

        renderCart(cart.items);

    } catch (error) {

        console.error(error);

        renderEmptyCart();
    }
}


// ======================================================
// PANIER VIDE
// ======================================================

function renderEmptyCart() {

    cartItemsContainer.innerHTML = `
        <div class="cart-empty">
            🛒
            <br><br>
            Votre panier est vide.
        </div>
    `;

    cartCount.textContent = "0";

    cartTotal.textContent = "0,00 €";

    checkoutButton.disabled = true;
    checkoutButton.textContent = "Panier vide";
}


// ======================================================
// AFFICHAGE DU PANIER
// ======================================================

function renderCart(items) {

    cartItemsContainer.innerHTML = "";

    let totalQuantity = 0;
    let totalAmount = 0;

    items.forEach(item => {

        const product =
            productsById[Number(item.product_id)];

        if (!product) {
            return;
        }

        const quantity =
            Number(item.quantity || 0);

        const price =
            Number(product.price || 0);

        totalQuantity += quantity;

        totalAmount +=
            price * quantity;

        const itemElement =
            document.createElement("div");

        itemElement.className = "cart-item";

        itemElement.innerHTML = `

            <div class="cart-item-image">
                ${getProductIllustration(product)}
            </div>

            <div>

                <div class="cart-item-name">
                    ${displayName(product)}
                </div>

                <div class="cart-item-price">
                    ${formatPrice(price)}
                </div>

                <div
                    style="
                        display:flex;
                        align-items:center;
                        gap:8px;
                        margin-top:10px;
                    "
                >

                    <button
                        type="button"
                        onclick="changeQuantity(
                            ${product.id},
                            ${quantity - 1}
                        )"
                        style="
                            border:none;
                            width:30px;
                            height:30px;
                            border-radius:6px;
                            font-size:18px;
                        "
                    >
                        −
                    </button>

                    <strong>
                        ${quantity}
                    </strong>

                    <button
                        type="button"
                        onclick="changeQuantity(
                            ${product.id},
                            ${quantity + 1}
                        )"
                        style="
                            border:none;
                            width:30px;
                            height:30px;
                            border-radius:6px;
                            font-size:18px;
                        "
                    >
                        +
                    </button>

                </div>

            </div>

            <button
                type="button"
                title="Supprimer le produit"
                onclick="removeCartItem(${product.id})"
                style="
                    border:none;
                    background:#fff1f1;
                    color:#dc3545;
                    width:36px;
                    height:36px;
                    border-radius:8px;
                    font-size:17px;
                "
            >
                🗑
            </button>
        `;

        cartItemsContainer.appendChild(
            itemElement
        );
    });


    // Bouton vider le panier
    const clearContainer =
        document.createElement("div");

    clearContainer.style.padding = "18px 0";

    clearContainer.innerHTML = `
        <button
            type="button"
            onclick="clearCart()"
            style="
                width:100%;
                border:1px solid #dc3545;
                background:white;
                color:#dc3545;
                padding:11px;
                border-radius:8px;
                font-weight:700;
            "
        >
            🗑 Vider le panier
        </button>
    `;

    cartItemsContainer.appendChild(
        clearContainer
    );

    cartCount.textContent =
        totalQuantity;

    cartTotal.textContent =
        formatPrice(totalAmount);

    checkoutButton.disabled = false;
    checkoutButton.textContent = "🔒 Passer la commande";
}


// ======================================================
// MODIFIER LA QUANTITÉ
// ======================================================

async function changeQuantity(
    productId,
    newQuantity
) {

    // Si quantité = 0 :
    // on supprime complètement le produit
    if (newQuantity <= 0) {
        await removeCartItem(productId);
        return;
    }

    try {

        const response =
            await fetch(`/api/cart/${userId}`);

        if (!response.ok) {
            throw new Error(
                "Impossible de lire le panier"
            );
        }

        const cart =
            await response.json();

        const updatedItems =
            cart.items.map(item => {

                if (
                    Number(item.product_id) ===
                    Number(productId)
                ) {

                    return {
                        product_id:
                            Number(item.product_id),

                        quantity:
                            Number(newQuantity)
                    };
                }

                return {
                    product_id:
                        Number(item.product_id),

                    quantity:
                        Number(item.quantity)
                };
            });


        const updateResponse =
            await fetch(`/api/cart/${userId}`, {

                method: "PUT",

                headers: {
                    "Content-Type":
                        "application/json"
                },

                body: JSON.stringify({
                    items: updatedItems
                })
            });


        if (!updateResponse.ok) {
            throw new Error(
                "Impossible de modifier le panier"
            );
        }

        await loadCart();

    } catch (error) {

        console.error(error);

        showToast(
            "Impossible de modifier le panier",
            true
        );
    }
}


// ======================================================
// SUPPRIMER UN ARTICLE
// ======================================================

async function removeCartItem(productId) {

    try {

        const response =
            await fetch(
                `/api/cart/${userId}?product_id=${productId}`,
                {
                    method: "DELETE"
                }
            );

        if (!response.ok) {
            throw new Error(
                "Impossible de supprimer le produit"
            );
        }

        showToast(
            "Produit supprimé du panier ✓"
        );

        await loadCart();

    } catch (error) {

        console.error(error);

        showToast(
            "Impossible de supprimer le produit",
            true
        );
    }
}


// ======================================================
// VIDER LE PANIER
// ======================================================

async function clearCart() {

    try {

        const response =
            await fetch(`/api/cart/${userId}`, {
                method: "DELETE"
            });

        if (!response.ok) {
            throw new Error(
                "Impossible de vider le panier"
            );
        }

        showToast("Panier vidé ✓");

        await loadCart();

    } catch (error) {

        console.error(error);

        showToast(
            "Impossible de vider le panier",
            true
        );
    }
}


// ======================================================
// CHECKOUT
// ======================================================

const paymentDialog =
    document.getElementById("payment-dialog");

const paymentForm =
    document.getElementById("payment-form");

const cancelPaymentButton =
    document.getElementById("cancel-payment-button");

const confirmPaymentButton =
    document.getElementById("confirm-payment-button");


function checkout() {

    if (!paymentDialog) {
        showToast(
            "Formulaire de paiement indisponible",
            true
        );

        return;
    }

    paymentDialog.showModal();
}


function closePaymentDialog() {

    paymentDialog.close();
}


async function processPayment(event) {

    event.preventDefault();

    const cardNumber =
        document
            .getElementById("card-number")
            .value
            .replace(/\s/g, "");

    const cardExpiry =
        document
            .getElementById("card-expiry")
            .value
            .trim();

    const cardCvv =
        document
            .getElementById("card-cvv")
            .value
            .trim();

    if (!/^\d{16}$/.test(cardNumber)) {

        showToast(
            "Le numéro de carte doit contenir 16 chiffres",
            true
        );

        return;
    }

    if (!/^\d{2}\/\d{2}$/.test(cardExpiry)) {

        showToast(
            "La date doit respecter le format MM/AA",
            true
        );

        return;
    }

    if (!/^\d{3}$/.test(cardCvv)) {

        showToast(
            "Le CVV doit contenir 3 chiffres",
            true
        );

        return;
    }

    const order = {
        user_id: userId,
        card_number: cardNumber,
        card_expiry: cardExpiry,
        card_cvv: cardCvv
    };

    confirmPaymentButton.disabled = true;
    confirmPaymentButton.textContent =
        "Paiement en cours...";

    try {

        const response =
            await fetch("/api/checkout", {

                method: "POST",

                headers: {
                    "Content-Type":
                        "application/json"
                },

                body: JSON.stringify(order)
            });


        if (!response.ok) {

            const text =
                await response.text();

            console.error(
                "Erreur checkout:",
                response.status,
                text
            );

            showToast(
                "Le paiement a été refusé",
                true
            );

            return;
        }


        const result =
            await response.json();


        const clearResponse =
            await fetch(
                `/api/cart/${userId}`,
                {
                    method: "DELETE"
                }
            );


        if (!clearResponse.ok) {

            console.warn(
                "Commande créée, mais panier non vidé"
            );
        }


        paymentForm.reset();
        closePaymentDialog();
        closeCart();

        await loadCart();


        showToast(
            `Paiement accepté ✓ Commande ${result.order_id}`
        );

    } catch (error) {

        console.error(error);

        showToast(
            "Erreur pendant le paiement",
            true
        );

    } finally {

        confirmPaymentButton.disabled = false;
        confirmPaymentButton.textContent =
            "Payer et commander";
    }
}


checkoutButton.addEventListener(
    "click",
    checkout
);


cancelPaymentButton.addEventListener(
    "click",
    closePaymentDialog
);


paymentForm.addEventListener(
    "submit",
    processPayment
);

// ======================================================
// INITIALISATION
// ======================================================

async function init() {

    await loadProducts();

    await loadCart();
}

init();
