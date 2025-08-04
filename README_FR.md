# Projet de Démonstration GitHub Webhook

Un projet complet pour apprendre et démontrer les fonctionnalités des GitHub Webhooks, incluant une implémentation de serveur Go et une documentation pédagogique détaillée.

## 🚀 Présentation du Projet

Ce projet fournit un serveur complet de réception de GitHub Webhooks, supportant :

- Réception et traitement de divers événements GitHub (Issues, Comments, Push, etc.)
- Vérification de signature HMAC-SHA256 pour assurer la sécurité
- Support pour le développement local et le déploiement sur la plateforme Render
- Logs détaillés et traitement d'événements
- Point de terminaison de vérification de santé

## 📁 Structure du Projet

```
├── README.md              # Documentation du projet
├── webhook.md             # Tutoriel complet sur les Webhooks
├── claude-code-intro.md   # Introduction à l'outil Claude Code
├── go.mod                 # Configuration du module Go (répertoire racine)
├── render.yaml            # Configuration de déploiement Render
└── webhook-demo/          # Implémentation du serveur Webhook
    ├── go.mod            # Configuration du module Go
    └── server.go         # Code principal du serveur
```

## 🛠️ Démarrage Rapide

### Prérequis

- Go 1.21 ou version supérieure
- Git

### Exécution Locale

1. **Cloner le projet**
   ```bash
   git clone <repository-url>
   cd agent-test
   ```

2. **Lancer le serveur**
   ```bash
   cd webhook-demo
   go run server.go
   ```

3. **Tester le serveur**
   ```bash
   curl http://localhost:8080/health
   ```

### Configuration du Secret Webhook (recommandé)

1. **Générer une clé**
   ```bash
   openssl rand -hex 20
   ```

2. **Définir la variable d'environnement**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 Déploiement sur Render

Le projet inclut un fichier de configuration `render.yaml` pour un déploiement en un clic sur la plateforme Render :

1. Connecter le dépôt GitHub à Render
2. Sélectionner Web Service
3. Render utilisera automatiquement la configuration dans `render.yaml`

Après le déploiement, il sera automatiquement configuré pour :
- Définir la variable d'environnement `WEBHOOK_SECRET`
- Configurer les vérifications de santé
- Activer le déploiement automatique

## 📚 Tutoriel d'Utilisation

Pour un tutoriel détaillé, veuillez consulter :
- [**webhook.md**](webhook.md) - Tutoriel complet de configuration et d'utilisation des GitHub Webhooks
- Inclut le tunneling réseau local, la configuration GitHub, les méthodes de test et autres étapes détaillées

## 🎯 Événements Supportés

Le serveur peut traiter les événements GitHub suivants :

- **ping** - Test de connexion Webhook
- **issues** - Création, édition, fermeture d'Issues, etc.
- **issue_comment** - Commentaires d'Issues (support de détection de commandes)
- **push** - Poussée de code
- D'autres événements sont également enregistrés

## 🔧 Points de Terminaison API

- `POST /webhook` - Point de terminaison de réception d'événements Webhook
- `GET /health` - Point de terminaison de vérification de santé

## 🔒 Fonctionnalités de Sécurité

- Vérification de signature HMAC-SHA256
- Support du Secret GitHub Webhook
- Enregistrement détaillé des logs de sécurité
- Validation d'entrée et gestion d'erreurs

## 📖 Documentation Associée

- [Introduction à Claude Code](claude-code-intro.md) - Introduction à l'outil assistant de programmation IA
- [Tutoriel Webhook](webhook.md) - Enseignement complet depuis zéro

## 🤝 Contribution

Les Issues et Pull Requests sont les bienvenues pour améliorer ce projet !

## 📄 Licence

Ce projet est destiné uniquement à des fins d'apprentissage et de démonstration.