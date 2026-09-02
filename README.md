# ShopSphere

ShopSphere est une application e-commerce microservices conçue comme support
d'un projet DevOps / DevSecOps / SRE.

L'objectif principal du projet n'est pas de développer une plateforme
e-commerce complète, mais de disposer d'une application réaliste permettant
de mettre en œuvre :

- conteneurisation
- CI/CD
- Kubernetes
- GitOps
- observabilité
- sécurité
- haute disponibilité
- sauvegarde et restauration
- pratiques SRE

---

## Architecture applicative

```text
                         Utilisateur
                              |
                              v
                     +----------------+
                     |    Frontend    |
                     |     Nginx      |
                     +-------+--------+
                             |
            +----------------+----------------+
            |                |                |
            v                v                v
     Product Catalog        Cart           Checkout
            |                |                |
            v                v          +-----+------+
         Percona           Valkey       |            |
                                        v            v
                                      Payment      Kafka
                                                     |
                                                     v
                                                Accounting
                                                     |
                                                     v
                                                  Percona


# Microservices
#Frontend
Interface web de ShopSphere.
Technologies :
HTML
JavaScript
Nginx
Nginx sert les fichiers statiques et joue également le rôle de reverse proxy
vers les APIs internes.

#Product Catalog
Gestion et consultation du catalogue produits.
API principale :
GET /products
Stockage :
Percona Server
Cart
Gestion du panier utilisateur.
APIs :
GET  /cart/{user_id}
POST /cart/{user_id}

Le panier supporte plusieurs produits et incrémente la quantité lorsqu'un
produit déjà présent est ajouté.
Stockage :

Valkey

# Checkout
Orchestre la création d'une commande.
Le service communique avec :

Cart
Product Catalog
Payment
Kafka


Flux simplifié :

Checkout
   |
   +--> Cart
   |
   +--> Product Catalog
   |
   +--> Payment
   |
   +--> Kafka


#Payment
Simulation d'un service de paiement.
API :

POST /payment

Le service retourne un identifiant de transaction et un statut.
Les informations de carte ne sont ni stockées dans Percona ni publiées dans
Kafka.


#Accounting
Service asynchrone consommant les commandes publiées dans Kafka.

Kafka
  |
  v
Accounting
  |
  v
Percona


Consumer Group Kafka :

accounting-service

Les commandes validées sont enregistrées dans la table orders.

#Services d'infrastructure

#Percona Server
Base de données relationnelle utilisée pour :
le catalogue produits
les commandes
Le schéma est initialisé automatiquement depuis :
db/init/01-init.sql


#Valkey
Stockage en mémoire utilisé par le service Cart.
#Kafka
Bus événementiel utilisé pour découpler Checkout et Accounting.
Topic :
orders

Flux :
Checkout --> Kafka --> Accounting

#Conteneurisation
Chaque microservice Go possède son propre Dockerfile.
Les services Go utilisent un build multi-stage :

golang builder
      |
      v
binary Go
      |
      v
distroless nonroot


L'image finale utilise :
gcr.io/distroless/static-debian12:nonroot

Cela permet notamment de réduire :
la taille des images
la surface d'attaque
les composants inutiles dans les conteneurs
Les services applicatifs sont exécutés avec un utilisateur non-root.

#Lancement local
#Prérequis
Docker
Docker Compose
Git

#Cloner le dépôt
git clone https://github.com/BA3S-AWS/shopsphere.git
cd shopsphere

#Démarrer l'application

docker compose up -d --build
#Vérifier les services
docker compose ps

Percona doit apparaître avec le statut :
healthy

#Accès à l'application
Frontend :
http://localhost:8080

Catalogue via Nginx :
http://localhost:8080/api/products


#Test du catalogue
curl http://localhost:8080/api/products

#Test du panier
Ajouter un produit :
curl -X POST http://localhost:8080/api/cart/demo-user \
  -H "Content-Type: application/json" \
  -d '{"items":[{"product_id":1,"quantity":1}]}'

Afficher le panier :
curl http://localhost:8080/api/cart/demo-user

#Test Checkout
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -d '{
    "user_id":"demo-user",
    "card_number":"4111111111111111",
    "card_expiry":"12/30",
    "card_cvv":"123"
  }'
Les informations de paiement utilisées ici sont uniquement des données
fictives pour le simulateur Payment du projet.

#Vérification Kafka / Accounting
Logs Accounting :
docker compose logs accounting
Une commande correctement consommée produit un message similaire à :
Order <UUID> saved successfully

#Vérification Percona
docker exec -it shopsphere-percona-1 \
  mysql -u root -prootpassword shopsphere \
  -e "SELECT order_id,user_id,total,status FROM orders;"

#Arrêt de l'environnement
Conserver les données :
docker compose down
Supprimer également les volumes :
docker compose down -v
Au prochain démarrage sur un volume vierge, le schéma Percona et les données
initiales seront automatiquement recréés.


#Évolution DevOps
Cette application servira ensuite de support à l'industrialisation :

Code
 |
 v
Git
 |
 v
CI/CD
 |
 v
Registry
 |
 v
Kubernetes
 |
 v
GitOps / ArgoCD
 |
 +--> Observabilité
 |
 +--> DevSecOps
 |
 +--> Autoscaling
 |
 +--> Backup / Restore

#Les prochaines étapes incluent notamment :
GitLab CI/CD
Kubernetes
Helm
ArgoCD
Cilium
Longhorn
Percona sur Kubernetes
Kafka sur Kubernetes
Prometheus
Grafana
Loki
Tempo / OpenTelemetry
Alertmanager
KEDA / HPA
gestion sécurisée des secrets
Velero / sauvegarde objet
tests de résilience

#Statut
MVP applicatif validé en environnement Docker Compose.
Flux end-to-end validé :

Frontend
   |
   v
Checkout
   |
   +--> Cart --> Valkey
   |
   +--> Product Catalog --> Percona
   |
   +--> Payment
   |
   +--> Kafka --> Accounting --> Percona



