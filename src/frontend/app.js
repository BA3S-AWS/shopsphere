const productsContainer = document.getElementById("products");
const cartContainer = document.getElementById("cart");
let userId = localStorage.getItem("shopsphere-user-id");

if (!userId) {
    userId = `user-${crypto.randomUUID()}`;
    localStorage.setItem("shopsphere-user-id", userId);
}

async function loadProducts() {
try {
const response = await fetch("/api/products");
    if (!response.ok) {
        throw new Error("Unable to load products");
    }

    const products = await response.json();

    productsContainer.innerHTML = "";

    products.forEach(product => {
        const div = document.createElement("div");

        div.innerHTML = `
            <h3>${product.name}</h3>
            <p>${product.description}</p>
            <p>Prix : ${product.price} €</p>
            <p>Stock : ${product.stock}</p>

            <button onclick="addToCart(${product.id})">
                Ajouter au panier
            </button>
        `;

        productsContainer.appendChild(div);
    });
} catch (error) {
    productsContainer.innerHTML =
        "<p>Impossible de charger le catalogue.</p>";
}
}
async function addToCart(productId) {
const cart = {
items: [
{
product_id: productId,
quantity: 1
}
]
};
const response = await fetch(`/api/cart/${userId}`, {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify(cart)
});

if (response.ok) {
    alert("Produit ajouté au panier");
    loadCart();
} else {
    alert("Impossible d'ajouter le produit au panier");
}
}
async function loadCart() {
try {
const response = await fetch(`/api/cart/${userId}`);
    if (!response.ok) {
        cartContainer.innerHTML = "<p>Panier vide.</p>";
        return;
    }

    const cart = await response.json();

    cartContainer.innerHTML = "";

    cart.items.forEach(item => {
        const div = document.createElement("div");

        div.innerHTML = `
            <p>Produit ID : ${item.product_id}</p>
            <p>Quantité : ${item.quantity}</p>
        `;

        cartContainer.appendChild(div);
    });
} catch (error) {
    cartContainer.innerHTML =
        "<p>Impossible de charger le panier.</p>";
}
}

async function checkout() {
    const checkoutButton = document.getElementById("checkout-button");

    const order = {
        user_id: userId,
        card_number: "4111111111111111",
        card_expiry: "12/30",
        card_cvv: "123"
    };

    checkoutButton.disabled = true;
    checkoutButton.textContent = "Traitement en cours...";

    try {
        const response = await fetch("/api/checkout", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(order)
        });

        if (!response.ok) {
            const message = await response.text();

            throw new Error(
                `Checkout ${response.status}: ${message}`
            );
        }

        const result = await response.json();

        // La commande est créée : vider le panier.
        const clearResponse = await fetch(
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

        cartContainer.innerHTML = `
            <section class="order-confirmation">
                <h2>✅ Commande confirmée</h2>

                <p>
                    Votre paiement de test a été accepté.
                </p>

                <p>
                    <strong>Commande :</strong>
                    ${result.order_id}
                </p>

                <p>
                    <strong>Montant :</strong>
                    ${Number(result.total).toFixed(2)} €
                </p>

                <p>
                    <strong>Statut :</strong>
                    ${result.status}
                </p>

                <p>
                    <strong>Livraison estimée :</strong>
                    ${result.estimated_days} jours
                </p>
            </section>
        `;

        checkoutButton.textContent = "Commande confirmée";
    } catch (error) {
        console.error(error);

        alert(
            "Impossible de finaliser la commande. " +
            "Veuillez réessayer."
        );

        checkoutButton.disabled = false;
        checkoutButton.textContent = "Commander";
    }
}

loadProducts();
loadCart();
