// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Artifact struct {
		Algorithm func(childComplexity int) int
		Digest    func(childComplexity int) int
	}

	Builder struct {
		URI func(childComplexity int) int
	}

	CVE struct {
		CveID func(childComplexity int) int
		Year  func(childComplexity int) int
	}

	CVEId struct {
		ID func(childComplexity int) int
	}

	GHSA struct {
		GhsaID func(childComplexity int) int
	}

	GHSAId struct {
		ID func(childComplexity int) int
	}

	HashEqual struct {
		Artifacts     func(childComplexity int) int
		Collector     func(childComplexity int) int
		Justification func(childComplexity int) int
		Origin        func(childComplexity int) int
	}

	IsOccurrence struct {
		Collector           func(childComplexity int) int
		Justification       func(childComplexity int) int
		OccurrenceArtifacts func(childComplexity int) int
		Origin              func(childComplexity int) int
		Package             func(childComplexity int) int
		Source              func(childComplexity int) int
	}

	OSV struct {
		OsvID func(childComplexity int) int
	}

	OSVId struct {
		ID func(childComplexity int) int
	}

	Package struct {
		Namespaces func(childComplexity int) int
		Type       func(childComplexity int) int
	}

	PackageName struct {
		Name     func(childComplexity int) int
		Versions func(childComplexity int) int
	}

	PackageNamespace struct {
		Names     func(childComplexity int) int
		Namespace func(childComplexity int) int
	}

	PackageQualifier struct {
		Key   func(childComplexity int) int
		Value func(childComplexity int) int
	}

	PackageVersion struct {
		Qualifiers func(childComplexity int) int
		Subpath    func(childComplexity int) int
		Version    func(childComplexity int) int
	}

	Query struct {
		Artifacts     func(childComplexity int, artifactSpec *model.ArtifactSpec) int
		Builders      func(childComplexity int, builderSpec *model.BuilderSpec) int
		Cve           func(childComplexity int, cveSpec *model.CVESpec) int
		Ghsa          func(childComplexity int, ghsaSpec *model.GHSASpec) int
		HashEquals    func(childComplexity int, hashEqualSpec *model.HashEqualSpec) int
		IsOccurrences func(childComplexity int, isOccurrenceSpec *model.IsOccurrenceSpec) int
		Osv           func(childComplexity int, osvSpec *model.OSVSpec) int
		Packages      func(childComplexity int, pkgSpec *model.PkgSpec) int
		Sources       func(childComplexity int, sourceSpec *model.SourceSpec) int
	}

	Source struct {
		Namespaces func(childComplexity int) int
		Type       func(childComplexity int) int
	}

	SourceName struct {
		Commit func(childComplexity int) int
		Name   func(childComplexity int) int
		Tag    func(childComplexity int) int
	}

	SourceNamespace struct {
		Names     func(childComplexity int) int
		Namespace func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Artifact.algorithm":
		if e.complexity.Artifact.Algorithm == nil {
			break
		}

		return e.complexity.Artifact.Algorithm(childComplexity), true

	case "Artifact.digest":
		if e.complexity.Artifact.Digest == nil {
			break
		}

		return e.complexity.Artifact.Digest(childComplexity), true

	case "Builder.uri":
		if e.complexity.Builder.URI == nil {
			break
		}

		return e.complexity.Builder.URI(childComplexity), true

	case "CVE.cveId":
		if e.complexity.CVE.CveID == nil {
			break
		}

		return e.complexity.CVE.CveID(childComplexity), true

	case "CVE.year":
		if e.complexity.CVE.Year == nil {
			break
		}

		return e.complexity.CVE.Year(childComplexity), true

	case "CVEId.id":
		if e.complexity.CVEId.ID == nil {
			break
		}

		return e.complexity.CVEId.ID(childComplexity), true

	case "GHSA.ghsaId":
		if e.complexity.GHSA.GhsaID == nil {
			break
		}

		return e.complexity.GHSA.GhsaID(childComplexity), true

	case "GHSAId.id":
		if e.complexity.GHSAId.ID == nil {
			break
		}

		return e.complexity.GHSAId.ID(childComplexity), true

	case "HashEqual.artifacts":
		if e.complexity.HashEqual.Artifacts == nil {
			break
		}

		return e.complexity.HashEqual.Artifacts(childComplexity), true

	case "HashEqual.collector":
		if e.complexity.HashEqual.Collector == nil {
			break
		}

		return e.complexity.HashEqual.Collector(childComplexity), true

	case "HashEqual.justification":
		if e.complexity.HashEqual.Justification == nil {
			break
		}

		return e.complexity.HashEqual.Justification(childComplexity), true

	case "HashEqual.origin":
		if e.complexity.HashEqual.Origin == nil {
			break
		}

		return e.complexity.HashEqual.Origin(childComplexity), true

	case "IsOccurrence.collector":
		if e.complexity.IsOccurrence.Collector == nil {
			break
		}

		return e.complexity.IsOccurrence.Collector(childComplexity), true

	case "IsOccurrence.justification":
		if e.complexity.IsOccurrence.Justification == nil {
			break
		}

		return e.complexity.IsOccurrence.Justification(childComplexity), true

	case "IsOccurrence.occurrenceArtifacts":
		if e.complexity.IsOccurrence.OccurrenceArtifacts == nil {
			break
		}

		return e.complexity.IsOccurrence.OccurrenceArtifacts(childComplexity), true

	case "IsOccurrence.origin":
		if e.complexity.IsOccurrence.Origin == nil {
			break
		}

		return e.complexity.IsOccurrence.Origin(childComplexity), true

	case "IsOccurrence.package":
		if e.complexity.IsOccurrence.Package == nil {
			break
		}

		return e.complexity.IsOccurrence.Package(childComplexity), true

	case "IsOccurrence.source":
		if e.complexity.IsOccurrence.Source == nil {
			break
		}

		return e.complexity.IsOccurrence.Source(childComplexity), true

	case "OSV.osvId":
		if e.complexity.OSV.OsvID == nil {
			break
		}

		return e.complexity.OSV.OsvID(childComplexity), true

	case "OSVId.id":
		if e.complexity.OSVId.ID == nil {
			break
		}

		return e.complexity.OSVId.ID(childComplexity), true

	case "Package.namespaces":
		if e.complexity.Package.Namespaces == nil {
			break
		}

		return e.complexity.Package.Namespaces(childComplexity), true

	case "Package.type":
		if e.complexity.Package.Type == nil {
			break
		}

		return e.complexity.Package.Type(childComplexity), true

	case "PackageName.name":
		if e.complexity.PackageName.Name == nil {
			break
		}

		return e.complexity.PackageName.Name(childComplexity), true

	case "PackageName.versions":
		if e.complexity.PackageName.Versions == nil {
			break
		}

		return e.complexity.PackageName.Versions(childComplexity), true

	case "PackageNamespace.names":
		if e.complexity.PackageNamespace.Names == nil {
			break
		}

		return e.complexity.PackageNamespace.Names(childComplexity), true

	case "PackageNamespace.namespace":
		if e.complexity.PackageNamespace.Namespace == nil {
			break
		}

		return e.complexity.PackageNamespace.Namespace(childComplexity), true

	case "PackageQualifier.key":
		if e.complexity.PackageQualifier.Key == nil {
			break
		}

		return e.complexity.PackageQualifier.Key(childComplexity), true

	case "PackageQualifier.value":
		if e.complexity.PackageQualifier.Value == nil {
			break
		}

		return e.complexity.PackageQualifier.Value(childComplexity), true

	case "PackageVersion.qualifiers":
		if e.complexity.PackageVersion.Qualifiers == nil {
			break
		}

		return e.complexity.PackageVersion.Qualifiers(childComplexity), true

	case "PackageVersion.subpath":
		if e.complexity.PackageVersion.Subpath == nil {
			break
		}

		return e.complexity.PackageVersion.Subpath(childComplexity), true

	case "PackageVersion.version":
		if e.complexity.PackageVersion.Version == nil {
			break
		}

		return e.complexity.PackageVersion.Version(childComplexity), true

	case "Query.artifacts":
		if e.complexity.Query.Artifacts == nil {
			break
		}

		args, err := ec.field_Query_artifacts_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Artifacts(childComplexity, args["artifactSpec"].(*model.ArtifactSpec)), true

	case "Query.builders":
		if e.complexity.Query.Builders == nil {
			break
		}

		args, err := ec.field_Query_builders_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Builders(childComplexity, args["builderSpec"].(*model.BuilderSpec)), true

	case "Query.cve":
		if e.complexity.Query.Cve == nil {
			break
		}

		args, err := ec.field_Query_cve_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Cve(childComplexity, args["cveSpec"].(*model.CVESpec)), true

	case "Query.ghsa":
		if e.complexity.Query.Ghsa == nil {
			break
		}

		args, err := ec.field_Query_ghsa_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Ghsa(childComplexity, args["ghsaSpec"].(*model.GHSASpec)), true

	case "Query.HashEquals":
		if e.complexity.Query.HashEquals == nil {
			break
		}

		args, err := ec.field_Query_HashEquals_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.HashEquals(childComplexity, args["hashEqualSpec"].(*model.HashEqualSpec)), true

	case "Query.IsOccurrences":
		if e.complexity.Query.IsOccurrences == nil {
			break
		}

		args, err := ec.field_Query_IsOccurrences_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.IsOccurrences(childComplexity, args["isOccurrenceSpec"].(*model.IsOccurrenceSpec)), true

	case "Query.osv":
		if e.complexity.Query.Osv == nil {
			break
		}

		args, err := ec.field_Query_osv_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Osv(childComplexity, args["osvSpec"].(*model.OSVSpec)), true

	case "Query.packages":
		if e.complexity.Query.Packages == nil {
			break
		}

		args, err := ec.field_Query_packages_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Packages(childComplexity, args["pkgSpec"].(*model.PkgSpec)), true

	case "Query.sources":
		if e.complexity.Query.Sources == nil {
			break
		}

		args, err := ec.field_Query_sources_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Sources(childComplexity, args["sourceSpec"].(*model.SourceSpec)), true

	case "Source.namespaces":
		if e.complexity.Source.Namespaces == nil {
			break
		}

		return e.complexity.Source.Namespaces(childComplexity), true

	case "Source.type":
		if e.complexity.Source.Type == nil {
			break
		}

		return e.complexity.Source.Type(childComplexity), true

	case "SourceName.commit":
		if e.complexity.SourceName.Commit == nil {
			break
		}

		return e.complexity.SourceName.Commit(childComplexity), true

	case "SourceName.name":
		if e.complexity.SourceName.Name == nil {
			break
		}

		return e.complexity.SourceName.Name(childComplexity), true

	case "SourceName.tag":
		if e.complexity.SourceName.Tag == nil {
			break
		}

		return e.complexity.SourceName.Tag(childComplexity), true

	case "SourceNamespace.names":
		if e.complexity.SourceNamespace.Names == nil {
			break
		}

		return e.complexity.SourceNamespace.Names(childComplexity), true

	case "SourceNamespace.namespace":
		if e.complexity.SourceNamespace.Namespace == nil {
			break
		}

		return e.complexity.SourceNamespace.Namespace(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputArtifactSpec,
		ec.unmarshalInputBuilderSpec,
		ec.unmarshalInputCVESpec,
		ec.unmarshalInputGHSASpec,
		ec.unmarshalInputHashEqualSpec,
		ec.unmarshalInputIsOccurrenceSpec,
		ec.unmarshalInputOSVSpec,
		ec.unmarshalInputPackageQualifierInput,
		ec.unmarshalInputPkgSpec,
		ec.unmarshalInputSourceQualifierInput,
		ec.unmarshalInputSourceSpec,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schema/artifact.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the artifact. It contains the algorithm and digest fields
"""
Artifact represents the artifact and contains a digest field

algorithm is mandatory in the from strings.ToLower(string(checksum.Algorithm)) (sha256, sha1...etc)
digest is mandatory in the form checksum.Value.

"""
type Artifact {
  algorithm: String!
  digest: String!
}

"""
ArtifactSpec allows filtering the list of artifacts to return.
"""
input ArtifactSpec {
  algorithm: String
  digest: String
}


extend type Query {
  "Returns all artifacts"
  artifacts(artifactSpec: ArtifactSpec): [Artifact!]!
}
`, BuiltIn: false},
	{Name: "../schema/builder.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the builder. It only contains the uri
"""
Builder represents the builder such as (FRSCA or github actions) and contains a uri field

uri is mandatory and represents the specific builder.

This node is a singleton: backends guarantee that there is exactly one node with
the same ` + "`" + `uri` + "`" + ` value.

"""
type Builder {
  uri: String!
}

"""
BuilderSpec allows filtering the list of builders to return.
"""
input BuilderSpec {
  uri: String
}


extend type Query {
  "Returns all builders"
  builders(builderSpec: BuilderSpec): [Builder!]!
}
`, BuiltIn: false},
	{Name: "../schema/cve.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the cve trie/tree. It contains the year
# along with the id associated with vulnerability (cve id)
"""
CVE represents common vulnerabilities and exposures. It contains the year along
with the CVE ID.

year is mandatory.

This node is a singleton: backends guarantee that there is exactly one node with
the same ` + "`" + `year` + "`" + ` value.

"""
type CVE {
  year: String!
  cveId: [CVEId!]!
}

"""
CVEId is the actual ID that is given to a specific vulnerability

id field is mandatory.

This node can be referred to by other parts of GUAC.
"""
type CVEId {
  id: String!
}

"""
CVESpec allows filtering the list of cves to return.
"""
input CVESpec {
  year: String
  cveId: String
}


extend type Query {
  "Returns all cve"
  cve(cveSpec: CVESpec): [CVE!]!
}
`, BuiltIn: false},
	{Name: "../schema/ghsa.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the ghsa trie/tree. It contains 
# id associated with vulnerability (ghsa id)
"""
GHSA represents github security advisory. It contains the ghsa ID (GHSA-pgvh-p3g4-86jw)

"""
type GHSA {
  ghsaId: [GHSAId!]!
}

"""
GHSAId is the actual ID that is given to a specific vulnerability on github

id field is mandatory.

This node can be referred to by other parts of GUAC.
"""
type GHSAId {
  id: String!
}

"""
GHSASpec allows filtering the list of ghsa to return.
"""
input GHSASpec {
  ghsaId: String
}


extend type Query {
  "Returns all ghsa"
  ghsa(ghsaSpec: GHSASpec): [GHSA!]!
}
`, BuiltIn: false},
	{Name: "../schema/hashEqual.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the HashEqual. It contains the justification, artifacts, source and collector. 
"""
HashEqual is an attestation represents when two artifact hash are similar based on a justification.

Justification - string value representing why the artifacts are the equal
Origin - where this attestation was generated from (based on which document)
Collector - the GUAC collector that collected the document that generated this attestation
Artifacts - the artifacts (represented by algorithm and digest) that are equal

"""
type HashEqual {
  justification: String!
  artifacts: [Artifact!]!
  origin: String!
  collector: String!
}

"""
HashEqualSpec allows filtering the list of HashEqual to return.

Specifying just the artifacts allows to query for all equivalent artifacts (if they exist)
"""
input HashEqualSpec {
  justification: String
  artifacts: [ArtifactSpec]
  origin: String
  collector: String
}


extend type Query {
  "Returns all HashEqual"
  HashEquals(hashEqualSpec: HashEqualSpec): [HashEqual!]!
}
`, BuiltIn: false},
	{Name: "../schema/isOccurrence.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the IsOccurrence. It contains the justification, package, source object,
#  occurrenceArtifacts, source of the attestation, and collector
"""
IsOccurrence is an attestation represents when either a package or source is represented by an artifact
Justification - string value representing why the package or source is represented by the specified artifact
Package - the package object type that represents the package
Source - the source object type that represents the source
occurrenceArtifacts - list of artifacts that represent the the package or source
Origin - where this attestation was generated from (based on which document)
Collector - the GUAC collector that collected the document that generated this attestation

Note: Package or Source must be specified but not both at the same time.
Attestation must occur at the PackageName or the PackageVersion or at the SourceName.

IsOccurrence does not connect a package with a source. 
HasSourceAt attestation will be used to connect a package with a source

"""
type IsOccurrence {
  justification: String!
  package: Package
  source: Source
  occurrenceArtifacts: [Artifact!]!
  origin: String!
  collector: String!
}

"""
IsOccurrenceSpec allows filtering the list of IsOccurrence to return.
Note: Package or Source must be specified but not both at the same time
For Package - a PackageName or PackageVersion must be specified (name or name, version, qualifiers and subpath)
Fro Source - a SourceName must be specified (name, tag or commit)
"""
input IsOccurrenceSpec {
  justification: String
  package: PkgSpec
  source: SourceSpec
  artifacts: [ArtifactSpec]
  origin: String
  collector: String
}


extend type Query {
  "Returns all IsOccurrence"
  IsOccurrences(isOccurrenceSpec: IsOccurrenceSpec): [IsOccurrence!]!
}`, BuiltIn: false},
	{Name: "../schema/osv.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the osv trie/tree. It contains 
# id associated with vulnerability (osv id)
"""
OSV represents Open Source Vulnerability . It contains a OSV ID.

"""
type OSV {
  osvId: [OSVId!]!
}

"""
OSVId is the actual ID that is given to a specific vulnerability

id field is mandatory. This maps to a GHSA or CVE ID

This node can be referred to by other parts of GUAC.
"""
type OSVId {
  id: String!
}

"""
OSVSpec allows filtering the list of osv to return.
"""
input OSVSpec {
  osvId: String
}


extend type Query {
  "Returns all osv"
  osv(osvSpec: OSVSpec): [OSV!]!
}
`, BuiltIn: false},
	{Name: "../schema/package.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the package trie/tree. This tree closely matches
# the pURL specification (https://github.com/package-url/purl-spec/blob/0dd92f26f8bb11956ffdf5e8acfcee71e8560407/README.rst)
# but deviates from it where GUAC rules state otherwise. In principle, we want
# this to represent a trie for packages, so information that represents a
# smaller collection of packages is being pushed downwards in the trie.

"""
Package represents a package.

In the pURL representation, each Package matches a ` + "`" + `pkg:<type>` + "`" + ` partial pURL.
The ` + "`" + `type` + "`" + ` field matches the pURL types but we might also use ` + "`" + `"guac"` + "`" + ` for the
cases where the pURL representation is not complete or when we have custom
rules.

This node is a singleton: backends guarantee that there is exactly one node with
the same ` + "`" + `type` + "`" + ` value.

Also note that this is named ` + "`" + `Package` + "`" + `, not ` + "`" + `PackageType` + "`" + `. This is only to make
queries more readable.
"""
type Package {
  type: String!
  namespaces: [PackageNamespace!]!
}

"""
PackageNamespace is a namespace for packages.

In the pURL representation, each PackageNamespace matches the
` + "`" + `pkg:<type>/<namespace>/` + "`" + ` partial pURL.

Namespaces are optional and type specific. Because they are optional, we use
empty string to denote missing namespaces.
"""
type PackageNamespace {
  namespace: String!
  names: [PackageName!]!
}

"""
PackageName is a name for packages.

In the pURL representation, each PackageName matches the
` + "`" + `pkg:<type>/<namespace>/<name>` + "`" + ` pURL.

Names are always mandatory.

This is the first node in the trie that can be referred to by other parts of
GUAC.
"""
type PackageName {
  name: String!
  versions: [PackageVersion!]!
}

"""
PackageVersion is a package version.

In the pURL representation, each PackageName matches the
` + "`" + `pkg:<type>/<namespace>/<name>@<version>` + "`" + ` pURL.

Versions are optional and each Package type defines own rules for handling them.
For this level of GUAC, these are just opaque strings.

This node can be referred to by other parts of GUAC.

Subpath and qualifiers are optional. Lack of qualifiers is represented by an
empty list and lack of subpath by empty string (to be consistent with
optionality of namespace and version). Two nodes that have different qualifiers
and/or subpath but the same version mean two different packages in the trie
(they are different). Two nodes that have same version but qualifiers of one are
a subset of the qualifier of the other also mean two different packages in the
trie.
"""
type PackageVersion {
  version: String!
  qualifiers: [PackageQualifier!]!
  subpath: String!
}

"""
PackageQualifier is a qualifier for a package, a key-value pair.

In the pURL representation, it is a part of the ` + "`" + `<qualifiers>` + "`" + ` part of the
` + "`" + `pkg:<type>/<namespace>/<name>@<version>?<qualifiers>` + "`" + ` pURL.

Qualifiers are optional, each Package type defines own rules for handling them,
and multiple qualifiers could be attached to the same package.

This node cannot be directly referred by other parts of GUAC.
"""
type PackageQualifier {
  key: String!
  value: String!
}

"""
PkgSpec allows filtering the list of packages to return.

Each field matches a qualifier from pURL. Use ` + "`" + `null` + "`" + ` to match on all values at
that level. For example, to get all packages in GUAC backend, use a PkgSpec
where every field is ` + "`" + `null` + "`" + `.

Empty string at a field means matching with the empty string. If passing in
qualifiers, all of the values in the list must match. Since we want to return
nodes with any number of qualifiers if no qualifiers are passed in the input, we
must also return the same set of nodes it the qualifiers list is empty. To match
on nodes that don't contain any qualifier, set ` + "`" + `matchOnlyEmptyQualifiers` + "`" + ` to
true. If this field is true, then the qualifiers argument is ignored.
"""
input PkgSpec {
  type: String
  namespace: String
  name: String
  version: String
  qualifiers: [PackageQualifierInput!] = []
  matchOnlyEmptyQualifiers: Boolean = false
  subpath: String
}

"""
PackageQualifierInput is the same as PackageQualifier, but usable as query
input.

GraphQL does not allow input types to contain composite types and does not allow
composite types to contain input types. So, although in this case these two
types are semantically the same, we have to duplicate the definition.

Keys are mandatory, but values could also be ` + "`" + `null` + "`" + ` if we want to match all
values for a specific key.

TODO(mihaimaruseac): Formalize empty vs null when the schema is fully done
"""
input PackageQualifierInput {
  key: String!
  value: String
}

extend type Query {
  "Returns all packages"
  packages(pkgSpec: PkgSpec): [Package!]!
}
`, BuiltIn: false},
	{Name: "../schema/source.graphql", Input: `#
# Copyright 2023 The GUAC Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# NOTE: This is experimental and might change in the future!

# Defines a GraphQL schema for the source trie/tree. This tree is a derivative of
# the pURL specification where it has a type, namespace, name and finally a qualifier that
# contain the tag or commit. 

"""
Source represents a source.

This can be the version control system that is being used.

This node is a singleton: backends guarantee that there is exactly one node with
the same ` + "`" + `type` + "`" + ` value.

Also note that this is named ` + "`" + `Source` + "`" + `, not ` + "`" + `SourceType` + "`" + `. This is only to make
queries more readable.
"""
type Source {
  type: String!
  namespaces: [SourceNamespace!]!
}

"""
SourceNamespace is a namespace for sources.

This can be represented as the location of the repo (such as github/gitlab/bitbucket)

Namespaces are optional and type specific. Because they are optional, we use
empty string to denote missing namespaces.
"""
type SourceNamespace {
  namespace: String!
  names: [SourceName!]!
}

"""
SourceName is a url of the repository and its tag or commit.

SourceName is mandatory. Either a tag or commit needs to be specified.

This is the first node in the trie that can be referred to by other parts of
GUAC.
"""
type SourceName {
  name: String!
  tag: String
  commit: String
}

"""
SourceSpec allows filtering the list of sources to return.

Empty string at a field means matching with the empty string.
"""
input SourceSpec {
  type: String
  namespace: String
  name: String
  qualifier: SourceQualifierInput
}

"""
SourceQualifierInput is the same as SourceQualifier, but usable as query
input.
"""
input SourceQualifierInput {
  tag: String
  commit: String
}

extend type Query {
  "Returns all sources"
  sources(sourceSpec: SourceSpec): [Source!]!
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
