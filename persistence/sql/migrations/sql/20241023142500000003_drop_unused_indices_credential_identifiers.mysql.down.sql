CREATE INDEX identity_credential_identifiers_nid_id_idx ON identity_credential_identifiers (nid ASC, id ASC);
CREATE INDEX identity_credential_identifiers_id_nid_idx ON identity_credential_identifiers (id ASC, nid ASC);
CREATE INDEX identity_credential_identifiers_nid_identity_credential_id_idx ON identity_credential_identifiers (identity_credential_id ASC, nid ASC);

DROP INDEX identity_credential_identifiers_identity_credential_id_idx ON identity_credential_identifiers;
