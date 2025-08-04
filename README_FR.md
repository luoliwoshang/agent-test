# Projet de Démonstration GitHub Webhook

Un projet complet pour apprendre et démontrer les fonctionnalités des GitHub Webhooks, incluant une implémentation de serveur Go et une documentation pédagogique détaillée.

## 🚀 Présentation du Projet

Ce projet fournit un serveur de réception GitHub Webhook complet qui prend en charge :

- Réception et traitement de divers événements GitHub (Issues, Comments, Push, etc.)
- Vérification de signature HMAC-SHA256 pour assurer la sécurité
- Support pour le développement local et le déploiement sur la plateforme Render
- Sortie de logs détaillée et traitement d'événements
- Point de terminaison de vérification de santé

## 📁 Structure du Projet

```
├── README.md              # Documentation du projet
├── webhook.md             # Tutoriel complet sur les Webhooks
├── claude-code-intro.md   # Introduction à l'outil Claude Code
├── go.mod                 # Configuration des modules Go (répertoire racine)
├── render.yaml            # Configuration de déploiement Render
└── webhook-demo/          # Implémentation du serveur Webhook
    ├── go.mod            # Configuration des modules Go
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

### Configuration du Secret Webhook (Recommandé)

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

Le projet inclut un fichier de configuration `render.yaml` permettant un déploiement en un clic sur la plateforme Render :

1. Connecter le dépôt GitHub à Render
2. Sélectionner Web Service
3. Render utilisera automatiquement la configuration dans `render.yaml`

Après le déploiement, cela configurera automatiquement :
- La variable d'environnement `WEBHOOK_SECRET`
- La vérification de santé
- Le déploiement automatique

## 📚 Tutoriel d'Utilisation

Pour un tutoriel d'utilisation détaillé, consultez :
- [**webhook.md**](webhook.md) - Tutoriel complet de configuration et d'utilisation des GitHub Webhooks
- Inclut des étapes détaillées pour le tunneling réseau, la configuration GitHub, les méthodes de test, etc.

## 🎯 Événements Pris en Charge

Le serveur peut traiter les événements GitHub suivants :

- **ping** - Test de connexion Webhook
- **issues** - Création, modification, fermeture d'Issues, etc.
- **issue_comment** - Commentaires d'Issues (avec détection de commandes)
- **push** - Poussée de code
- D'autres événements sont également enregistrés

## 🔧 Points d'API

- `POST /webhook` - Point de terminaison de réception d'événements Webhook
- `GET /health` - Point de terminaison de vérification de santé

## 🔒 Fonctionnalités de Sécurité

- Vérification de signature HMAC-SHA256
- Support du Secret Webhook GitHub
- Enregistrement détaillé des journaux de sécurité
- Validation d'entrée et gestion d'erreurs

## 📖 Documentation Connexe

- [Introduction à Claude Code](claude-code-intro.md) - Présentation de l'outil assistant de programmation IA
- [Tutoriel Webhook](webhook.md) - Enseignement complet depuis zéro

## 🤝 Contribution

Les Issues et Pull Requests sont les bienvenues pour améliorer ce projet !

## 📄 Licence

Ce projet est uniquement destiné à des fins d'apprentissage et de démonstration.