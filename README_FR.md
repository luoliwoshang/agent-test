# Projet de Démonstration GitHub Webhook

Un projet complet pour apprendre et démontrer les fonctionnalités des GitHub Webhooks, incluant une implémentation serveur Go et une documentation tutorielle détaillée.

## 🚀 Aperçu du Projet

Ce projet fournit un serveur complet de réception de GitHub Webhooks qui prend en charge :

- Réception et traitement de divers événements GitHub (Issues, Commentaires, Push, etc.)
- Vérification de signature HMAC-SHA256 pour la sécurité
- Support pour le développement local et le déploiement sur la plateforme Render
- Journalisation détaillée et traitement des événements
- Point de terminaison de vérification de santé

## 📁 Structure du Projet

```
├── README.md              # Documentation du projet
├── webhook.md             # Tutoriel complet sur les Webhooks
├── claude-code-intro.md   # Introduction à l'outil Claude Code
├── go.mod                 # Configuration du module Go (racine)
├── render.yaml            # Configuration de déploiement Render
└── webhook-demo/          # Implémentation du serveur Webhook
    ├── go.mod            # Configuration du module Go
    └── server.go         # Code du serveur principal
```

## 🛠️ Démarrage Rapide

### Prérequis

- Go 1.21 ou version supérieure
- Git

### Développement Local

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

### Configurer le Secret Webhook (Recommandé)

1. **Générer un secret**
   ```bash
   openssl rand -hex 20
   ```

2. **Définir la variable d'environnement**
   ```bash
   export WEBHOOK_SECRET=your-generated-secret
   cd webhook-demo
   go run server.go
   ```

## 🌐 Déployer sur Render

Le projet inclut un fichier de configuration `render.yaml` pour un déploiement en un clic sur la plateforme Render :

1. Connecter le dépôt GitHub à Render
2. Sélectionner Web Service
3. Render utilisera automatiquement la configuration dans `render.yaml`

Après le déploiement, cela configurera automatiquement :
- La variable d'environnement `WEBHOOK_SECRET`
- Les vérifications de santé
- Le déploiement automatique

## 📚 Tutoriel d'Utilisation

Pour des tutoriels détaillés d'utilisation, veuillez consulter :
- [**webhook.md**](webhook.md) - Tutoriel complet de configuration et d'utilisation des GitHub Webhooks
- Inclut des étapes détaillées pour le tunneling, la configuration GitHub, les méthodes de test, etc.

## 🎯 Événements Pris en Charge

Le serveur peut traiter les événements GitHub suivants :

- **ping** - Test de connexion Webhook
- **issues** - Création, édition, fermeture d'Issues, etc.
- **issue_comment** - Commentaires d'Issues (prend en charge la détection de commandes)
- **push** - Poussées de code
- D'autres événements sont également enregistrés

## 🔧 Points de Terminaison API

- `POST /webhook` - Point de terminaison de réception d'événements Webhook
- `GET /health` - Point de terminaison de vérification de santé

## 🔒 Fonctionnalités de Sécurité

- Vérification de signature HMAC-SHA256
- Support pour GitHub Webhook Secret
- Journalisation de sécurité détaillée
- Validation d'entrée et gestion d'erreurs

## 📖 Documentation Associée

- [Introduction à Claude Code](claude-code-intro.md) - Introduction à l'outil assistant de programmation IA
- [Tutoriel Webhook](webhook.md) - Tutoriel complet depuis zéro

## 🤝 Contribution

N'hésitez pas à soumettre des Issues et des Pull Requests pour améliorer ce projet !

## 📄 Licence

Ce projet est uniquement à des fins d'apprentissage et de démonstration.